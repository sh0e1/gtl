package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sh0e1/gtl/translate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	apiKey       string
	gcpProjectID string

	source string
	target string
)

var rootCmd = &cobra.Command{
	Use:   "gtl [text to translate]",
	Short: "Translate input text",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if tmp := os.Getenv("GTL_API_KEY"); tmp == "" && apiKey == "" {
			return errors.New("api-key is empty")
		} else if tmp != "" && apiKey == "" {
			apiKey = tmp
		}

		if tmp := os.Getenv("GTL_GCP_PROJECT_ID"); tmp == "" && gcpProjectID == "" {
			return errors.New("gcp-project-id is empty")
		} else if tmp != "" && gcpProjectID == "" {
			gcpProjectID = tmp
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		c, err := translate.New(ctx, gcpProjectID, apiKey)
		if err != nil {
			return err
		}
		defer c.Close()

		translated, err := c.TranslateText(ctx, source, target, args)
		if err != nil {
			return nil
		}
		for _, t := range translated {
			cmd.Println(t.GetTranslatedText())
		}
		return nil
	},
	TraverseChildren: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOut(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func init() {
	const (
		defaultTarget = "ja"
		defaultSource = "en"
	)

	cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gtl.yaml)")
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "default is environment variable $GTL_API_KEY")
	rootCmd.PersistentFlags().StringVar(&gcpProjectID, "gcp-project-id", "", "default is $GTL_GCP_PROJECT_ID")

	rootCmd.Flags().StringVar(&source, "source", defaultSource, "BCP-47 language code of input text")
	rootCmd.Flags().StringVar(&target, "target", defaultTarget, "BCP-47 language code of used to translate input text")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gtl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gtl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

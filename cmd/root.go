package cmd

import (
	"context"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sh0e1/gtl/translate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	apiKey    string
	projectID string

	target string
	source string
)

const (
	defaultTarget = "ja"
	defaultSource = "en"
)

var rootCmd = &cobra.Command{
	Use: "gtl",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		c, err := translate.New(ctx, projectID, apiKey)
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gtl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", os.Getenv("GTL_API_KEY"), "")
	rootCmd.PersistentFlags().StringVarP(&projectID, "gcp-project", "p", os.Getenv("GTL_GCP_PROJECT_ID"), "")

	rootCmd.Flags().StringVarP(&target, "target", "t", defaultTarget, "")
	rootCmd.Flags().StringVarP(&source, "source", "s", defaultSource, "")
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

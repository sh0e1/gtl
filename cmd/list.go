package cmd

import (
	"context"

	"github.com/sh0e1/gtl/translate"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use: "list",
	//Short: "",
	//Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		c, err := translate.New(ctx, projectID, apiKey)
		if err != nil {
			return err
		}
		defer c.Close()

		langs, err := c.GetSupportedLanguages(ctx, language.Japanese)
		if err != nil {
			return err
		}
		for _, l := range langs {
			cmd.Printf("%-5s : %s\n", l.GetLanguageCode(), l.GetDisplayName())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

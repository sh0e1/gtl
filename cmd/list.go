package cmd

import (
	"context"

	"github.com/sh0e1/gtl/translate"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays a list of supported languages",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := rootCmd.PersistentFlags().GetString("api-key")
		gcpProjectID, _ := rootCmd.PersistentFlags().GetString("gcp-project-id")

		ctx := context.Background()
		c, err := translate.New(ctx, gcpProjectID, apiKey)
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
}

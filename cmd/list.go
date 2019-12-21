package cmd

import (
	"context"

	translate "cloud.google.com/go/translate/apiv3"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	"google.golang.org/grpc/metadata"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use: "list",
	//Short: "",
	//Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := metadata.AppendToOutgoingContext(context.Background(), "x-goog-api-key", apiKey)
		ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-user-project", projectID)
		client, err := translate.NewTranslationClient(ctx)
		if err != nil {
			return err
		}
		defer client.Close()

		req := &translatepb.GetSupportedLanguagesRequest{
			Parent:              "projects/" + projectID,
			DisplayLanguageCode: language.Japanese.String(),
		}
		resp, err := client.GetSupportedLanguages(ctx, req)
		if err != nil {
			return err
		}

		for _, l := range resp.GetLanguages() {
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

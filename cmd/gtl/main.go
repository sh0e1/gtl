package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	"google.golang.org/grpc/metadata"
)

const (
	defaultTarget = "ja"
	defaultSource = "en"
)

var (
	apiKey    = os.Getenv("GTL_API_KEY")
	projectID = os.Getenv("GTL_GCP_PROJECT_ID")
)

func main() {
	var (
		target = flag.String("target", defaultTarget, "")
		source = flag.String("source", defaultSource, "")
	)
	flag.Parse()

	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-goog-api-key", apiKey)
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-user-project", projectID)
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Contents:           flag.Args(),
		SourceLanguageCode: *source,
		TargetLanguageCode: *target,
		Parent:             "projects/" + projectID,
	}
	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range resp.GetTranslations() {
		fmt.Println(t.GetTranslatedText())
	}
}

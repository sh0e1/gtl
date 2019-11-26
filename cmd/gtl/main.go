package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func main() {
	const (
		defaultTarget = "ja"
		defaultSource = "en"
	)

	var (
		target = flag.String("target", defaultTarget, "")
		source = flag.String("source", defaultSource, "")
	)
	flag.Parse()

	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(os.Getenv("GTL_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	targetLang, err := language.Parse(*target)
	if err != nil {
		log.Fatal(err)
	}

	sourceLang, err := language.Parse(*source)
	if err != nil {
		log.Fatal(err)
	}

	opt := &translate.Options{
		Source: sourceLang,
	}
	translations, err := client.Translate(ctx, flag.Args(), targetLang, opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range translations {
		fmt.Println(t.Text)
	}
}

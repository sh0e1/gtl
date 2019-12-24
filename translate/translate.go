package translate

import (
	"context"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	"google.golang.org/grpc/metadata"
)

// NewClient ...
func NewClient(ctx context.Context, projectID, apiKey string) (*Client, error) {
	c, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		TranslationClient: c,
		projectID:         projectID,
		apiKey:            apiKey,
	}, nil
}

// Client ...
type Client struct {
	*translate.TranslationClient
	projectID string
	apiKey    string
}

// Close ...
func (c *Client) Close() {
	c.TranslationClient.Close()
}

// TranslateText ...
func (c *Client) TranslateText(ctx context.Context, source, target string, contents []string) ([]string, error) {
	req := &translatepb.TranslateTextRequest{
		Contents:           contents,
		SourceLanguageCode: source,
		TargetLanguageCode: target,
		Parent:             "projects/" + c.projectID,
	}
	resp, err := c.TranslationClient.TranslateText(c.appendContext(ctx), req)
	if err != nil {
		return nil, err
	}

	translated := make([]string, 0, len(resp.GetTranslations()))
	for _, t := range resp.GetTranslations() {
		translated = append(translated, t.GetTranslatedText())
	}
	return translated, nil
}

func (c *Client) appendContext(ctx context.Context) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-user-project", c.projectID)
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-api-key", c.apiKey)
	return ctx
}

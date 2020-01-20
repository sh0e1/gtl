package translate

import (
	"context"

	translate "cloud.google.com/go/translate/apiv3"
	"golang.org/x/text/language"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	"google.golang.org/grpc/metadata"
)

// Translator is translator client interface.
type Translator interface {
	TranslateText(ctx context.Context, source, target string, contents []string) ([]*translatepb.Translation, error)
	GetSupportedLanguages(ctx context.Context, lang language.Tag) ([]*translatepb.SupportedLanguage, error)
	Close()
}

// New returns translator client interface.
func New(ctx context.Context, projectID, apiKey string) (Translator, error) {
	c, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		TranslationClient: c,
		ProjectID:         projectID,
		ApiKey:            apiKey,
	}, nil
}

// Client
type Client struct {
	*translate.TranslationClient
	ProjectID string
	ApiKey    string
}

// Close closes the connection.
func (c *Client) Close() {
	_ = c.TranslationClient.Close()
}

// TranslateText translates text.
func (c *Client) TranslateText(ctx context.Context, source, target string, contents []string) ([]*translatepb.Translation, error) {
	req := &translatepb.TranslateTextRequest{
		Contents:           contents,
		SourceLanguageCode: source,
		TargetLanguageCode: target,
		Parent:             c.parent(),
	}
	resp, err := c.TranslationClient.TranslateText(c.appendToOutgoingContext(ctx), req)
	if err != nil {
		return nil, err
	}
	return resp.GetTranslations(), nil
}

// GetSupportedLanguages gets support languages.
func (c *Client) GetSupportedLanguages(ctx context.Context, lang language.Tag) ([]*translatepb.SupportedLanguage, error) {
	req := &translatepb.GetSupportedLanguagesRequest{
		Parent:              c.parent(),
		DisplayLanguageCode: lang.String(),
	}
	resp, err := c.TranslationClient.GetSupportedLanguages(c.appendToOutgoingContext(ctx), req)
	if err != nil {
		return nil, err
	}
	return resp.GetLanguages(), nil
}

func (c *Client) appendToOutgoingContext(ctx context.Context) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-user-project", c.ProjectID)
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-api-key", c.ApiKey)
	return ctx
}

func (c *Client) parent() string {
	return "projects/" + c.ProjectID
}

package translate

import (
	"context"

	translate "cloud.google.com/go/translate/apiv3"
	"golang.org/x/text/language"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	"google.golang.org/grpc/metadata"
)

// Translator ...
type Translator interface {
	TranslateText(ctx context.Context, source, target string, contents []string) ([]*translatepb.Translation, error)
	GetSupportedLanguages(ctx context.Context, lang language.Tag) ([]*translatepb.SupportedLanguage, error)
	Close()
}

// New ...
func New(ctx context.Context, projectID, apiKey string) (Translator, error) {
	c, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return nil, err
	}
	return &client{
		TranslationClient: c,
		projectID:         projectID,
		apiKey:            apiKey,
	}, nil
}

type client struct {
	*translate.TranslationClient
	projectID string
	apiKey    string
}

// Close ...
func (c *client) Close() {
	_ = c.TranslationClient.Close()
}

// TranslateText ...
func (c *client) TranslateText(ctx context.Context, source, target string, contents []string) ([]*translatepb.Translation, error) {
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

// GetSupportedLanguages ...
func (c *client) GetSupportedLanguages(ctx context.Context, lang language.Tag) ([]*translatepb.SupportedLanguage, error) {
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

func (c *client) appendToOutgoingContext(ctx context.Context) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-user-project", c.projectID)
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-api-key", c.apiKey)
	return ctx
}

func (c *client) parent() string {
	return "projects/" + c.projectID
}

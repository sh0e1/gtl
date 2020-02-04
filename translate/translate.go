package translate

import (
	"context"

	translate "cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

// Translator is translator client interface.
type Translator interface {
	Translate(ctx context.Context, source, target string, contents []string) ([]translate.Translation, error)
	GetSupportedLanguages(ctx context.Context, target language.Tag) ([]translate.Language, error)
	Close()
}

// New returns translator client interface.
func New(ctx context.Context, apiKey string) (Translator, error) {
	c, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: c,
	}, nil
}

// Client
type Client struct {
	*translate.Client
}

// Close closes the connection.
func (c *Client) Close() {
	_ = c.Client.Close()
}

// Translate translates text.
func (c *Client) Translate(ctx context.Context, source, target string, contents []string) ([]translate.Translation, error) {
	s, err := language.Parse(source)
	if err != nil {
		return nil, err
	}
	t, err := language.Parse(target)
	if err != nil {
		return nil, err
	}

	opt := &translate.Options{
		Source: s,
	}
	ts, err := c.Client.Translate(ctx, contents, t, opt)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// GetSupportedLanguages gets support languages.
func (c *Client) GetSupportedLanguages(ctx context.Context, target language.Tag) ([]translate.Language, error) {
	return c.Client.SupportedLanguages(ctx, target)
}

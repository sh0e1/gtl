package translate_test

import (
	"context"
	"errors"
	"os"
	"testing"

	googletranslate "cloud.google.com/go/translate"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sh0e1/gtl/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func TestNew(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	apiKey := "api-key"

	tc, err := googletranslate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tc.Close()
	}()

	want := &translate.Client{
		Client: tc,
	}

	got, err := translate.New(ctx, apiKey)
	if err != nil {
		t.Errorf("err should be nil, but got %q", err)
		return
	}
	defer got.Close()

	if diff := cmp.Diff(want, got,
		cmpopts.IgnoreTypes(googletranslate.Client{})); diff != "" {
		t.Errorf("differs (-want +got):\n%s", diff)
	}
}

func TestClient_Translate(t *testing.T) {
	apiKey := os.Getenv("GTL_API_KEY")
	if apiKey == "" {
		t.Fatal("Must be set API Key to environment variable: GTL_API_KEY")
	}

	ctx := context.Background()
	c, err := translate.New(ctx, apiKey)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	tests := []struct {
		name     string
		source   string
		target   string
		contexts []string
		want     []googletranslate.Translation
		wantErr  error
	}{
		{
			name:     "Success Japanese to English",
			source:   language.Japanese.String(),
			target:   language.English.String(),
			contexts: []string{"こんにちは"},
			want: []googletranslate.Translation{
				{
					Text: "Hello",
				},
			},
			wantErr: nil,
		},
		{
			name:     "Success English to Japanese",
			source:   language.English.String(),
			target:   language.Japanese.String(),
			contexts: []string{"Hello"},
			want: []googletranslate.Translation{
				{
					Text: "こんにちは",
				},
			},
			wantErr: nil,
		},
		{
			name:     "Invalid source language",
			source:   "invalid",
			target:   language.Japanese.String(),
			contexts: []string{"Hello"},
			want:     nil,
			wantErr:  errors.New("language: tag is not well-formed"),
		},
		{
			name:     "Invalid target language",
			source:   language.English.String(),
			target:   "invalid",
			contexts: []string{"Hello"},
			want:     nil,
			wantErr:  errors.New("language: tag is not well-formed"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := c.Translate(context.Background(), tt.source, tt.target, tt.contexts)
			if tt.wantErr != nil {
				if err == nil {
					t.Error("expected error, but got nil")
					return
				}
				if g, e := err.Error(), tt.wantErr.Error(); g != e {
					t.Errorf("unexpected error:\nwant: %q\ngot : %q", e, g)
					return
				}
				return
			}

			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreTypes(language.Tag{})); diff != "" {
				t.Errorf("differs (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_GetSupportedLanguages(t *testing.T) {
	t.Parallel()
	t.Log("TODO: implement")
}

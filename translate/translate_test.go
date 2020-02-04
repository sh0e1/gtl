package translate_test

import (
	"context"
	"testing"

	googletranslate "cloud.google.com/go/translate"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sh0e1/gtl/translate"
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
	t.Parallel()
	t.Log("TODO: implement")
}

func TestClient_GetSupportedLanguages(t *testing.T) {
	t.Parallel()
	t.Log("TODO: implement")
}

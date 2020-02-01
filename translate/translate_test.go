package translate_test

import (
	"context"
	"testing"

	translatev3 "cloud.google.com/go/translate/apiv3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sh0e1/gtl/translate"
)

func TestNew(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	projectID := "project-id"
	apiKey := "api-key"

	tc, err := translatev3.NewTranslationClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	want := &translate.Client{
		TranslationClient: tc,
		ProjectID:         projectID,
		ApiKey:            apiKey,
	}

	got, err := translate.New(context.Background(), projectID, apiKey)
	if err != nil {
		t.Errorf("err should be nil, but got %q", err)
		return
	}
	defer got.Close()

	if diff := cmp.Diff(want, got,
		cmpopts.IgnoreTypes(translatev3.TranslationClient{})); diff != "" {
		t.Errorf("differs (-want +got):\n%s", diff)
	}
}

func TestClient_TranslateText(t *testing.T) {
	t.Parallel()
	t.Log("TODO: implement")
}

func TestClient_GetSupportedLanguages(t *testing.T) {
	t.Parallel()
	t.Log("TODO: implement")
}

package urls

import (
	"context"
	"net/http"
	"testing"

	"github.com/apex/log"
)

func TestIntegrationSuccess(t *testing.T) {
	config := Config{
		BaseURL:           "https://orchestrate.ooni.io",
		CountryCode:       "IT",
		EnabledCategories: []string{"NEWS", "CULTR"},
		HTTPClient:        http.DefaultClient,
		Limit:             17,
		Logger:            log.Log,
		UserAgent:         "ooniprobe-engine/v0.1.0-dev",
	}
	ctx := context.Background()
	result, err := Query(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Results) < 1 {
		t.Fatal("no results")
	}
}

func TestIntegrationFailure(t *testing.T) {
	config := Config{
		BaseURL:           "\t\t\t",
		CountryCode:       "IT",
		EnabledCategories: []string{"NEWS", "CULTR"},
		HTTPClient:        http.DefaultClient,
		Limit:             17,
		Logger:            log.Log,
		UserAgent:         "ooniprobe-engine/v0.1.0-dev",
	}
	ctx := context.Background()
	result, err := Query(ctx, config)
	if err == nil {
		t.Fatal("expected an error here")
	}
	if result != nil {
		t.Fatal("expected nil result here")
	}
}

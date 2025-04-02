package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	apiKey, err := auth.GetAPIKey(headers)

	if !reflect.DeepEqual("my-secret-key", apiKey) {
		t.Fatalf("expected: %v, got: %v", "my-secret-key", apiKey)
	}
	if !reflect.DeepEqual(nil, err) {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
}

func TestGetAPIKey_NoHeader(t *testing.T) {
	headers := http.Header{}

	apiKey, err := auth.GetAPIKey(headers)

	if !reflect.DeepEqual("", apiKey) {
		t.Fatalf("expected: %v, got: %v", "", apiKey)
	}

	if !reflect.DeepEqual("no authorization header included", err.Error()) {
		t.Fatalf("expected: %v, got: %v", "no authorization header included", err)
	}
}

/*
func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-secret-key")

	apiKey, err := auth.GetAPIKey(headers)

	assert.Error(t, err)
	assert.Equal(t, "malformed authorization header", err.Error())
	assert.Empty(t, apiKey)
}

func TestGetAPIKey_MissingKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey")

	apiKey, err := auth.GetAPIKey(headers)

	assert.Error(t, err)
	assert.Equal(t, "malformed authorization header", err.Error())
	assert.Empty(t, apiKey)
}*/

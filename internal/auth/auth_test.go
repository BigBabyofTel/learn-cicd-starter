package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key-123")

	apiKey, err := GetAPIKey(headers)
	require.NoError(t, err)
	require.Equal(t, "test-api-key-123", apiKey)
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no authorization header included")
}

func TestGetAPIKeyInvalidFormat(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer token-123")

	_, err := GetAPIKey(headers)
	require.Error(t, err)
	require.Contains(t, err.Error(), "malformed authorization header")
}

func TestGetAPIKeyEmptyKey(t *testing.T) {
	headers := http.Header{}
	// not validating apikey
	headers.Set("Authorization", " ")

	_, err := GetAPIKey(headers)
	require.Error(t, err)
	require.Contains(t, err.Error(), "malformed authorization header")
}

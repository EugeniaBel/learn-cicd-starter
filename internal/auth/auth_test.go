package auth

import (
	"net/http"
	"testing"
)

// Test 1: Check that valid API key is extracted correctly
func TestGetAPIKey_ValidKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key-123")

	result, err := GetAPIKey(headers)

	// Should not return error
	if err != nil {
		t.Errorf("Got error: %v, want no error", err)
	}

	// Should return correct key
	if result != "test-key-123" {
		t.Errorf("Got: %s, want: test-key-123", result)
	}
}

// Test 2: Check that error is returned when header is missing
func TestGetAPIKey_NoHeader(t *testing.T) {
	headers := http.Header{} // Empty headers

	_, err := GetAPIKey(headers)

	// Should return error
	if err == nil {
		t.Errorf("Should return error when no header, but got nil")
	}
}

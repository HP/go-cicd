package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key-123"},
			},
			expectedKey:   "test-api-key-123",
			expectedError: "",
		},
		{
			name:          "no authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "empty authorization header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "malformed header - no space",
			headers: http.Header{
				"Authorization": []string{"ApiKeytest-api-key-123"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "malformed header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-api-key-123"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "malformed header - only prefix",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "malformed header - only space",
			headers: http.Header{
				"Authorization": []string{" "},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "valid API key with extra spaces",
			headers: http.Header{
				"Authorization": []string{"ApiKey  test-api-key-123"},
			},
			expectedKey:   "",
			expectedError: "",
		},
		{
			name: "valid API key with multiple parts",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key-123 extra-part"},
			},
			expectedKey:   "test-api-key-123",
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() key = %v, expectedKey %v", key, tt.expectedKey)
			}

			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("GetAPIKey() error = %v, expectedError nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("GetAPIKey() error = nil, expectedError %v", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("GetAPIKey() error = %v, expectedError %v", err.Error(), tt.expectedError)
				}
			}
		})
	}
}

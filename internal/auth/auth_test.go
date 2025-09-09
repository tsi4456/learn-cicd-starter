package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	req, _ := http.NewRequest("GET", "blah", nil)
	baseHeader := req.Header
	missingAuthorisation := baseHeader.Clone()
	hasAuthorisation := baseHeader.Clone()
	hasAuthorisation.Add("Authorization", "ApiKey blahblahblah")
	shortAuth := hasAuthorisation.Clone()
	shortAuth.Set("Authorization", "ApiKeyHasNoSpaces")
	wrongAuth := hasAuthorisation.Clone()
	wrongAuth.Set("Authorization", "Ap1Key blahblahblahblah")

	tests := []struct {
		name     string
		input    http.Header
		want_str string
	}{
		{name: "no headers", input: missingAuthorisation, want_str: ""},
		{name: "correct header format", input: hasAuthorisation, want_str: "blahblahblah"},
		{name: "short header", input: shortAuth, want_str: ""},
		{name: "wrong header", input: wrongAuth, want_str: ""},
	}

	for _, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want_str, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want_str, got)
		}
	}
}

package practice

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/hello", "Hello, /hello"},
		{"/world", "Hello, /world"},
	}
	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, tt.path, nil)
		rec := httptest.NewRecorder()
		Hello(rec, req)
		if got := rec.Body.String(); got != tt.want {
			t.Errorf("Hello(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

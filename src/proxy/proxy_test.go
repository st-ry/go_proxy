package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	
	rootHandler(response, request)
	t.Run("returns hello", func(t *testing.T) {
		got := response.Body.String()
		want := "hello"
		
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
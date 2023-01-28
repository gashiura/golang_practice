package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserShowAction(t *testing.T) {
	tests := []struct {
		name string
		userID string
		want int
	}{
		{ "userID is integer", "100", 200 },
		{ "userID is string", "abcde", 404 },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := "/users/" + tt.userID
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			Routes().ServeHTTP(w, req)

			if w.Result().StatusCode != tt.want {
				t.Errorf("TestInfoRequest() test returned an unexpected result: got %v want %v", w.Result().StatusCode, tt.want)
			}
		})
	}
}

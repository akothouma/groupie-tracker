package handler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie/vars"
)

func TestGetArtists(t *testing.T) {
	tmpl, err := template.ParseGlob("../" + vars.Template_dir + "/*.html")
	if err != nil {
		t.Errorf("Failed to parse templates: %v", err)
	}
	vars.Templates = tmpl

	tests := []struct {
		name                 string
		method               string
		expected_status_code int
	}{
		{
			name:                 "GET with valid JSON",
			method:               http.MethodGet,
			expected_status_code: http.StatusOK,
		},
		{
			name:                 "non-GET method",
			method:               http.MethodPost,
			expected_status_code: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/", nil)

			w := httptest.NewRecorder()
			GetArtists(w, request)

			if w.Code != tc.expected_status_code {
				t.Errorf("expected %v, got %v", tc.expected_status_code, w.Code)
			}
		})
	}
}

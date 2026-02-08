package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Home Handler",
			handler:        home,
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Shapes API",
		},
		{
			name:           "Health Handler",
			handler:        health,
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is running",
		},
		{
			name:           "About Handler",
			handler:        about,
			expectedStatus: http.StatusOK,
			expectedBody:   "Alex Milian",
		},
		{
			name:           "Time Handler",
			handler:        currentTime,
			expectedStatus: http.StatusOK,
			expectedBody:   "", // time changes, just check not empty
		},
		{
			name:           "Greeting Handler",
			handler:        greeting,
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, welcome to the Shapes API",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("got status %v, expected %v", rr.Code, tt.expectedStatus)
			}

			if tt.expectedBody != "" && !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("got body %q, expected to contain %q", rr.Body.String(), tt.expectedBody)
			}

			if tt.name == "Time Handler" && rr.Body.String() == "" {
				t.Errorf("expected time response, got empty string")
			}
		})
	}
}

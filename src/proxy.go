package main

import (
	"net/http"
)

type MITMSettings struct {
	CookieReplacementRules map[string]string
}

func MITMHandler(settings MITMSettings) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// Check cookie replacement rules

		// ...

		// Forward the request

		// Recieve response

		// Write response

		// Log transaction

	}
}

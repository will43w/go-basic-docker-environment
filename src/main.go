package main

import (
	"log"
	"net/http"
)

// CLI for:
//   - setting ssl certificate files

// Create proxy based on CLI arguments

// Run proxy based on configurations

func main() {
	// Grab certificate files given command line arguments
	cert := ""
	key := ""

	// Set up proxies settings
	settings := MITMSettings{CookieReplacementRules: make(map[string]string)}

	// Define MITM implementation
	http.HandleFunc("/", MITMHandler(settings))

	// Start server
	log.Printf("About to listen on port 443")
	err := http.ListenAndServeTLS(":443", cert, key, nil)
	log.Fatal(err)

	// Enable settings editing through shell
}

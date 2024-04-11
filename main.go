package main

import (
	"fmt"
	"net/http"
)

// Middleware to set Referrer-Policy header
func setReferrerPolicyHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Referrer-Policy header to "strict-origin-when-cross-origin"
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a file server handler for serving static files from the "frontend" directory
	fs := http.FileServer(http.Dir("frontend"))

	// Use the custom middleware to set the Referrer-Policy header
	http.Handle("/", setReferrerPolicyHeader(fs))

	// Print message before starting the server
	fmt.Println("Listening on port 8080")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}




















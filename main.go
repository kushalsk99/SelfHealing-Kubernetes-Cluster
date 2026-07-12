package main

import (
	"fmt"
	"net/http"
)

// ❌ INTENTIONAL SECURITY FLAW FOR PIPELINE TESTING
// SonarQube will flag this plaintext hardcoded key immediately as a vulnerability
const MOCK_AWS_SECRET_KEY = "AKIAIOSFODNN7EXAMPLE/FAKEKEYDOEXTERNALISETHIS"

func main() {
	// A simple endpoint our cluster metrics and chaos engines can probe
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Healthy and listening")
	})

	fmt.Println("Starting target chaos-app service on port 8080...")
	fmt.Printf("Baking internal system configurations identifier: %s\n", MOCK_AWS_SECRET_KEY)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

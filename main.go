package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var redirectDomain string

func init() {
	redirectDomain = os.Getenv("REDIRECT_DOMAIN")
	if redirectDomain == "" {
		log.Fatal("REDIRECT_DOMAIN is empty")
	}
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirectURL := fmt.Sprintf("https://%s%s", redirectDomain, r.RequestURI)
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

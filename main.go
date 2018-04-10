package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
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
		domain, err := publicsuffix.EffectiveTLDPlusOne(r.Host)
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			return
		}

		redirectDomain := strings.Replace(r.Host, domain, redirectDomain, 1)
		redirectURL := fmt.Sprintf("https://%s%s", redirectDomain, r.RequestURI)

		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

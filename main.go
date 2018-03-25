package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// encodePath will encode the path portion of the URL
// this is against the http spec, but is the way AWS
// handles things
// encodePath will always return a leading slash

var cloudFrontURL string

func encodePath(originalPath string) string {
	tokens := strings.Split(originalPath, "/")
	var encodedTokens []string

	for _, element := range tokens {
		// QueryEscape not path escape because this violates the spec
		encodedTokens = append(encodedTokens, url.QueryEscape(element))
	}

	return strings.Join(encodedTokens, "/")
}

// redirect will URL encode the path and redirect to the other webserver
func redirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	encodedPath := encodePath(path)
	newURL := cloudFrontURL + encodedPath
	fmt.Printf("%s -> %s\n", path, newURL)
	http.Redirect(w, r, newURL, 302)
}

func main() {
	host := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	cloudFrontURL = os.Getenv("CLOUDFRONT_URL")

	listenOn := net.JoinHostPort(host, port)
	if cloudFrontURL == "" {
		log.Fatal("You must specify the CLOUDFRONT_URL (environment var)")
	}

	fmt.Printf("AWS Madness Corrector redirecting %s to %s\n", listenOn, cloudFrontURL)

	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(listenOn, nil)

	if err != nil {
		log.Fatal("Can't start: ", err)
	}
}

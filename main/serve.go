package main

import (
	"net/url"
	"flag"
	"net/http/httputil"
	"log"
	"net/http"
)

var uiVersion string

func main() {
	apiURLStr := flag.String("api", "http://localhost:9200", "the url to proxy API from")
	portStr := flag.String("port", "8080", "port")
	flag.Parse()

	apiURL, err := url.Parse(*apiURLStr)
	if err != nil {
		log.Fatal(err)
	}
	rProxy := httputil.NewSingleHostReverseProxy(apiURL)
	rProxy.ModifyResponse = func(r *http.Response) error {
		r.Header.Set("Kibanator-UI-Version", uiVersion)
		return nil
	}
	http.Handle("/", rProxy)
	http.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(assetFS())))
	//http.Handle("/", http.RedirectHandler("/ui/", http.StatusMovedPermanently))
	log.Fatal(http.ListenAndServe(":" + *portStr, nil))
}
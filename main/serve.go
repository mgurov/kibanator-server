package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var uiVersion = strings.TrimSpace((string)(MustAsset("ui/REACT_APP_VERSION")))

func main() {
	apiURLStr := flag.String("api", "", "the url to proxy API from")
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
	http.Handle("/api/", rProxy)
	http.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(assetFS())))
	//http.Handle("/", http.RedirectHandler("/ui/", http.StatusMovedPermanently))
	log.Println("ready at port", *portStr)
	log.Fatal(http.ListenAndServe(":"+*portStr, nil))
}

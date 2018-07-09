package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
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
	http.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(
		defaultToIndexHtmlFileSystem{assetFS()},
	)))
	//http.Handle("/", http.RedirectHandler("/ui/", http.StatusMovedPermanently))
	log.Println("ready at port", *portStr)
	log.Fatal(http.ListenAndServe(":"+*portStr, nil))
}

type defaultToIndexHtmlFileSystem struct {
	fs http.FileSystem
}

func (nfs defaultToIndexHtmlFileSystem) Open(path string) (http.File, error) {
	log.Println("Opening", path)
	f, err := nfs.fs.Open(path)
	if err != nil {
		log.Println("Error opening", err)
		if err == os.ErrNotExist {
			return nfs.fs.Open("/index.html")
		}
		return nil, err
	}

	return f, nil
}

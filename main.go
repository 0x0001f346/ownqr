package main

import (
	_ "embed"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed assets/index.html
var embededIndexHTML []byte

//go:embed assets/favicon.svg
var embededFaviconSVG []byte

func main() {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	router.GET("/", httpRoot)
	router.GET("/favicon.ico", httpFaviconICO)
	router.GET("/favicon.svg", httpFaviconSVG)
	router.GET("/qr/:errorCorrection/*body", httpQR)
	router.GET("/download/:errorCorrection/*body", httpQRDownload)

	err := http.ListenAndServe(":9100", router)
	if err != nil {
		panic(err)
	}
}

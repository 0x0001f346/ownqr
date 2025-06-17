package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func httpFaviconICO(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/favicon.svg", http.StatusSeeOther)
}

func httpFaviconSVG(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(embededFaviconSVG)
}

func httpQR(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body := strings.TrimPrefix(ps.ByName("body"), "/")
	errorCorrection := ps.ByName("errorCorrection")

	qr, err := createQRCodeAsPNG(
		body,
		errorCorrection,
	)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", getQRCodeFilename()))
	w.Write(qr)
}

func httpQRDownload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body := strings.TrimPrefix(ps.ByName("body"), "/")
	errorCorrection := ps.ByName("errorCorrection")

	qr, err := createQRCodeAsPNG(
		body,
		errorCorrection,
	)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", getQRCodeFilename()))
	w.Write(qr)
}

func httpRoot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(embededIndexHTML)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"
)

func generateQRcode(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Please provide a text parameter", http.StatusBadRequest)
		return
	}

	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}
	pngData, err := qr.PNG(256)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(pngData)))
	w.Write(pngData)
}

func main() {
	http.HandleFunc("/generate", generateQRcode)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

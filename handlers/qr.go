package handlers

import (
	"net/http"

	"github.com/emaforlin/CodeSquare/services"
)

func GetQRCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	data := r.URL.Query().Get("data")
	if data == "" || len(data) > 256 {
		http.Error(w, "invalid data received", http.StatusBadRequest)
		return
	}

	qrcode, err := services.CreateQRCode(data)
	if err != nil {
		http.Error(w, "failed to create QR code", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(qrcode)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/bos/internal/models/product"
)

func (a *App) ProductHandler(w http.ResponseWriter, r *http.Request) {

	barcode := r.URL.Query().Get("barcode")

	var response struct {
		Product product.Product `json:"product"`
		Status  string          `json:"status"`
	}

	for _, product := range a.ProductList {
		if product.ProductID == barcode {
			response.Product = product
			response.Status = "success"
			break
		} else {
			response.Status = "not found"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

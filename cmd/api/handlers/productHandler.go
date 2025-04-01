package handlers

import (
	"encoding/json"
	"net/http"
)

func (a *App) ProductHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a.ProductList)

}

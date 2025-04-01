package router

import (
	"net/http"

	"github.com/redmejia/bos/cmd/api/handlers"
)

func Router(app *handlers.App) http.Handler {

	mux := http.NewServeMux()

	var fs = http.FileServer(http.Dir("assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/api/v1/products", app.ProductHandler)

	return mux
}

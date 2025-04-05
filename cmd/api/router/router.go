package router

import (
	"net/http"

	"github.com/redmejia/bos/cmd/api/handlers"
	"github.com/redmejia/bos/cmd/api/middleware"
)

func Router(app *handlers.App) http.Handler {

	mux := http.NewServeMux()

	var fs = http.FileServer(http.Dir("assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/api/v1/products", middleware.IsAuthorized(app, app.ProductsHandler))
	mux.HandleFunc("/api/v1/product", middleware.IsAuthorized(app, app.ProductHandler))

	return middleware.Logger(mux)
}

package handlers

import (
	"net/http"
)

func (a *App) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	Render(a.Template, w, "index.html", map[string]interface{}{
		"products": a.ProductList,
	})
}

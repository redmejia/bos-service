package handlers

import (
	"net/http"

	"github.com/redmejia/bos/internal/models/product"
)

func (a *App) ProductsHandler(w http.ResponseWriter, r *http.Request) {

	var productList []product.Product

	for _, item := range a.ProductList {
		var formatCurrency = item.Price / 100.0
		productList = append(productList, product.Product{
			ProductID:    item.ProductID,
			Name:         item.Name,
			Price:        formatCurrency,
			ProductImage: item.ProductImage,
			BarcodeImage: item.BarcodeImage,
		})
	}

	Render(a.Template, w, "index.html", map[string]interface{}{
		"products": productList,
	})
}

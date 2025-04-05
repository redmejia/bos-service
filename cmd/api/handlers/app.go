package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/redmejia/bos/internal/models/product"
)

type App struct {
	Host           string
	Port           string
	Info, ErrorLog *log.Logger
	JWTKey         string
	// Issuer         string
	// Wg             *sync.WaitGroup
	ProductList []product.Product // This is a list of products, right now only in memory but this could be a database
	Template    *template.Template
}

func Render(t *template.Template, w http.ResponseWriter, name string, data interface{}) error {
	err := t.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}

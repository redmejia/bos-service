package handlers

import (
	"log"

	"github.com/redmejia/bos/internal/models/product"
)

type App struct {
	Host           string
	Port           string
	Info, ErrorLog *log.Logger
	// Wg             *sync.WaitGroup
	ProductList []product.Product // This is a list of products, right now only in memory but this could be a database
}

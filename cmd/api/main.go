package main

import (
	"sync"

	"github.com/redmejia/bos/internal/models/product"
	"github.com/redmejia/bos/internal/utils/barcode"
)

func main() {

	var wg sync.WaitGroup

	productList := []product.Product{
		{ProductID: "PROD-001", Name: "Product 1", Price: "10.00"},
		{ProductID: "PROD-002", Name: "Product 2", Price: "20.00"},
		{ProductID: "PROD-003", Name: "Product 3", Price: "30.00"},
		{ProductID: "PROD-004", Name: "Product 4", Price: "40.00"},
		{ProductID: "PROD-005", Name: "Product 5", Price: "50.00"},
		{ProductID: "PROD-006", Name: "Product 6", Price: "60.00"},
		{ProductID: "PROD-007", Name: "Product 7", Price: "70.00"},
		{ProductID: "PROD-008", Name: "Product 8", Price: "80.00"},
		{ProductID: "PROD-009", Name: "Product 9", Price: "90.00"},
		{ProductID: "PROD-010", Name: "Product 10", Price: "100.00"},
		{ProductID: "PROD-011", Name: "Product 11", Price: "110.00"},
		{ProductID: "PROD-012", Name: "Product 12", Price: "120.00"},
		{ProductID: "PROD-013", Name: "Product 13", Price: "130.00"},
		{ProductID: "PROD-014", Name: "Product 14", Price: "140.00"},
		{ProductID: "PROD-015", Name: "Product 15", Price: "150.00"},
		{ProductID: "PROD-016", Name: "Product 16", Price: "160.00"},
		{ProductID: "PROD-017", Name: "Product 17", Price: "170.00"},
		{ProductID: "PROD-018", Name: "Product 18", Price: "180.00"},
		{ProductID: "PROD-019", Name: "Product 19", Price: "190.00"},
		{ProductID: "PROD-020", Name: "Product 20", Price: "200.00"},
	}

	for _, product := range productList {
		wg.Add(1)
		go barcode.GenerateBarcodeList(&wg, product)
	}
	wg.Wait()

}

package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/redmejia/bos/cmd/api/router"
	"github.com/redmejia/bos/internal/models/product"
	"github.com/redmejia/bos/internal/utils/barcode"
)

func main() {

	var wg sync.WaitGroup

	productList := []product.Product{
		{ProductID: "PROD-001", Name: "Product 1", Price: "10.00", BarcodeImage: "http://localhost:8080/assets/PROD-001.png"},
		{ProductID: "PROD-002", Name: "Product 2", Price: "20.00", BarcodeImage: "http://localhost:8080/assets/PROD-002.png"},
		{ProductID: "PROD-003", Name: "Product 3", Price: "30.00", BarcodeImage: "http://localhost:8080/assets/PROD-003.png"},
		{ProductID: "PROD-004", Name: "Product 4", Price: "40.00", BarcodeImage: "http://localhost:8080/assets/PROD-004.png"},
		{ProductID: "PROD-005", Name: "Product 5", Price: "50.00", BarcodeImage: "http://localhost:8080/assets/PROD-005.png"},
		{ProductID: "PROD-006", Name: "Product 6", Price: "60.00", BarcodeImage: "http://localhost:8080/assets/PROD-006.png"},
		{ProductID: "PROD-007", Name: "Product 7", Price: "70.00", BarcodeImage: "http://localhost:8080/assets/PROD-007.png"},
		{ProductID: "PROD-008", Name: "Product 8", Price: "80.00", BarcodeImage: "http://localhost:8080/assets/PROD-008.png"},
		{ProductID: "PROD-009", Name: "Product 9", Price: "90.00", BarcodeImage: "http://localhost:8080/assets/PROD-009.png"},
		{ProductID: "PROD-010", Name: "Product 10", Price: "100.00", BarcodeImage: "http://localhost:8080/assets/PROD-010.png"},
		{ProductID: "PROD-011", Name: "Product 11", Price: "110.00", BarcodeImage: "http://localhost:8080/assets/PROD-011.png"},
		{ProductID: "PROD-012", Name: "Product 12", Price: "120.00", BarcodeImage: "http://localhost:8080/assets/PROD-012.png"},
		{ProductID: "PROD-013", Name: "Product 13", Price: "130.00", BarcodeImage: "http://localhost:8080/assets/PROD-013.png"},
		{ProductID: "PROD-014", Name: "Product 14", Price: "140.00", BarcodeImage: "http://localhost:8080/assets/PROD-014.png"},
		{ProductID: "PROD-015", Name: "Product 15", Price: "150.00", BarcodeImage: "http://localhost:8080/assets/PROD-015.png"},
		{ProductID: "PROD-016", Name: "Product 16", Price: "160.00", BarcodeImage: "http://localhost:8080/assets/PROD-016.png"},
		{ProductID: "PROD-017", Name: "Product 17", Price: "170.00", BarcodeImage: "http://localhost:8080/assets/PROD-017.png"},
		{ProductID: "PROD-018", Name: "Product 18", Price: "180.00", BarcodeImage: "http://localhost:8080/assets/PROD-018.png"},
		{ProductID: "PROD-019", Name: "Product 19", Price: "190.00", BarcodeImage: "http://localhost:8080/assets/PROD-019.png"},
		{ProductID: "PROD-020", Name: "Product 20", Price: "200.00", BarcodeImage: "http://localhost:8080/assets/PROD-020.png"},
	}

	for _, product := range productList {
		wg.Add(1)
		go barcode.GenerateBarcodeList(&wg, product)
	}
	wg.Wait()

	app := &router.App{
		Port:     ":8080",
		Info:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	srv := &http.Server{
		Addr:     app.Port,
		ErrorLog: app.ErrorLog,
		Handler:  router.Router(app),
	}

	app.Info.Printf("Starting server on %s", app.Port)

	if err := srv.ListenAndServe(); err != nil {
		app.ErrorLog.Fatalf("Error starting server: %s", err)
	}

}

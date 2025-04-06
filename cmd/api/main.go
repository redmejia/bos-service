package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"text/template"

	"github.com/redmejia/bos/cmd/api/handlers"
	"github.com/redmejia/bos/cmd/api/router"
	"github.com/redmejia/bos/internal/models/product"
	"github.com/redmejia/bos/internal/utils/barcode"
	"github.com/redmejia/bos/internal/utils/security/jwt"
)

func main() {

	var (
		port   string
		host   string
		jwtKey string
		issuer string
	)
	defaultPort := "8080"
	defaultHost := "localhost"
	flag.StringVar(&port, "port", defaultPort, "server port")
	flag.StringVar(&host, "host", defaultHost, "hostname")
	flag.StringVar(&jwtKey, "key", "", "JWT key")
	flag.StringVar(&issuer, "iss", "", "Issuer")
	flag.Parse()

	var wg sync.WaitGroup

	productList := []product.Product{
		{ProductID: "PROD-001", Name: "Product 1", Price: 0.60 * 100, ProductImage: fmt.Sprintf("http://%s:8080/assets/products/PROD-001.png", host), BarcodeImage: fmt.Sprintf("http://%s:8080/assets/PROD-001.png", host)},
		{ProductID: "PROD-002", Name: "Product 2", Price: 7.12 * 100, ProductImage: fmt.Sprintf("http://%s:8080/assets/products/PROD-002.png", host), BarcodeImage: fmt.Sprintf("http://%s:8080/assets/PROD-002.png", host)},
		{ProductID: "PROD-003", Name: "Product 3", Price: 10.50 * 100, ProductImage: fmt.Sprintf("http://%s:8080/assets/products/PROD-003.png", host), BarcodeImage: fmt.Sprintf("http://%s:8080/assets/PROD-003.png", host)},
		{ProductID: "PROD-004", Name: "Product 4", Price: 2.30 * 100, ProductImage: fmt.Sprintf("http://%s:8080/assets/products/PROD-004.png", host), BarcodeImage: fmt.Sprintf("http://%s:8080/assets/PROD-004.png", host)},
		{ProductID: "PROD-005", Name: "Product 5", Price: 5.30 * 100, ProductImage: fmt.Sprintf("http://%s:8080/assets/products/PROD-005.png", host), BarcodeImage: fmt.Sprintf("http://%s:8080/assets/PROD-005.png", host)},
		{ProductID: "PROD-006", Name: "Product 6", Price: 0.55 * 100, ProductImage: fmt.Sprintf("http://%s:8080/assets/products/PROD-006.png", host), BarcodeImage: fmt.Sprintf("http://%s:8080/assets/PROD-006.png", host)},
		// ... add more here
	}

	for _, product := range productList {
		wg.Add(1)
		go barcode.GenerateBarcodeList(&wg, product)
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("Barcodes generated successfully")

	wg.Wait()

	token, err := jwt.GenerateToken(jwtKey, issuer)

	if err != nil {
		errorLog.Fatalf("Error generating token: %s", err)

	}

	infoLog.Println("Token: ", token)

	tmpl := template.Must(template.ParseGlob("views/*.html"))

	app := &handlers.App{
		Port:     fmt.Sprintf(":%s", port),
		Info:     infoLog,
		ErrorLog: errorLog,
		JWTKey:   jwtKey,
		// Issuer:      issuer,
		ProductList: productList,
		Template:    tmpl,
	}

	srv := &http.Server{
		Addr:     app.Port,
		ErrorLog: app.ErrorLog,
		Handler:  router.Router(app),
	}

	infoLog.Printf("Starting server on %s\n", app.Port)

	if err := srv.ListenAndServe(); err != nil {
		app.ErrorLog.Fatalf("Error starting server: %s", err)
	}

}

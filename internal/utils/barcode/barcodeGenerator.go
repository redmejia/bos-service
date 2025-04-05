package barcode

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"sync"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/redmejia/bos/internal/models/product"
)

func GenerateBarcodeList(wg *sync.WaitGroup, product product.Product) {

	defer wg.Done()

	barcodeCode, err := code128.Encode(product.ProductID)
	if err != nil {
		fmt.Println("Error encoding barcode:", err)
		return
	}

	scaledBarcode, err := barcode.Scale(barcodeCode, 200, 80)
	if err != nil {
		fmt.Println("Error scaling barcode:", err)
		return
	}

	assetsDir := "assets"

	err = os.MkdirAll(assetsDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	barcodeCodeName := fmt.Sprintf("%s.png", product.ProductID)

	filepath := filepath.Join(assetsDir, barcodeCodeName)

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, scaledBarcode)
	if err != nil {
		fmt.Println("Error encoding PNG image:", err)
		return
	}

	// fmt.Printf("Barcode for %s generated successfully: %s\n", product.ProductID, barcodeCodeName)

}

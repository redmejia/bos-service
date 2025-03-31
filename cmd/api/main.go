package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func main() {
	// Test data
	barcodeCode, err := code128.Encode("PROD-002")
	if err != nil {
		fmt.Println("Error encoding barcode:", err)
		return
	}

	scaledBarcode, err := barcode.Scale(barcodeCode, 200, 80)
	if err != nil {
		fmt.Println("Error scaling barcode:", err)
	}

	assetsDir := "assets"

	err = os.MkdirAll(assetsDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filepath := filepath.Join(assetsDir, "barcode2.png")

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	err = png.Encode(file, scaledBarcode)
	if err != nil {
		fmt.Println("Error encoding PNG image:", err)
	}

	fmt.Println("Barcode encoded and saved to barcode.png")
}

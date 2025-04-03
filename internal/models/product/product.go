package product

type Product struct {
	ProductID    string  `json:"product_id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	ProductImage string  `json:"product_image"`
	BarcodeImage string  `json:"barcode_image"`
}

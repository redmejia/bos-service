package product

type Product struct {
	ProductID    string `json:"product_id"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	BarcodeImage string `json:"barcode_image"`
}

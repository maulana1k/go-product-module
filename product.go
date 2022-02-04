package product

type Product struct {
	name, brand, variant string
	stock, price         int
}

var ProductStorage []Product

func GetProduct() []Product {
	return ProductStorage
}

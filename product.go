package product

type Product struct {
	name, brand  string
	stock, price int
}

var ProductStorage = make([]Product, 0, 20)

func GetProduct() []Product {
	var product1 = Product{"laptop", "asus", 12, 8000000}
	ProductStorage = append(ProductStorage, product1)

	return ProductStorage
}

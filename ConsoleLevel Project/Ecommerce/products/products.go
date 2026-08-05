package products

var ProductList = make(map[int]Products, 4)

type Products struct {
	ProdId       int
	ProdName     string
	prodCategory string
	Quantity     int
	Price        int
}

var arr = []Products{
	{ProdId: 100, ProdName: "MRF Bat", prodCategory: "Sports", Quantity: 10, Price: 6000},
	{ProdId: 101, ProdName: "El Paso", prodCategory: "Perfumes", Quantity: 25, Price: 299},
	{ProdId: 170, ProdName: "Iphone 17 pro", prodCategory: "Accessories", Quantity: 6, Price: 80000},
	{ProdId: 192, ProdName: "Chess Board", prodCategory: "Sports", Quantity: 10, Price: 1000},
}

func (p Products) GetproductByID(prodId int) int {
	for items := range ProductList {
		if ProductList[items].ProdId == prodId {
			return items
		}
	}
	return -1
}

func (p Products) Addproducts() {
	for i := 1; i <= 4; i++ {
		ProductList[i] = arr[i-1]
	}
}

package payment

import (
	"ConsoleLevel_Project/Ecommerce/order"
	"ConsoleLevel_Project/Ecommerce/products"
	"fmt"
	"sync"
	"time"
)

func Payment(prodId int, quantity int, price int, wg *sync.WaitGroup) {
	prod := products.Products{}
	item := prod.GetproductByID(prodId)

	defer wg.Done()

	if item == -1 {
		fmt.Println("Product Not Found")
	}

	if products.ProductList[item].Price*quantity == price && products.ProductList[item].Quantity >= quantity {
		time.Sleep(3 * time.Second)
		fmt.Println("Payment Successful")
		fmt.Println(order.Orderproduct(item, wg))
	} else {
		time.Sleep(3 * time.Second)
		fmt.Println("Payment Failed")
	}
}

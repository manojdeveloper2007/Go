package order

import (
	"ConsoleLevel_Project/Ecommerce/products"
	"fmt"
	"sync"
	"time"
)

var p = products.Products{}

func Orderproduct(prod int, wg *sync.WaitGroup) string {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	res := fmt.Sprintf("Your order is confirmed\nProduct ID : %v\nProduct Name : %v", products.ProductList[prod].ProdId, products.ProductList[prod].ProdName)
	return res
}

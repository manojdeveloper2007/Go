package main

import (
	"ConsoleLevel_Project/Ecommerce/payment"
	"ConsoleLevel_Project/Ecommerce/products"
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	products.Products{}.Addproducts()
	product := products.ProductList
	fmt.Println("#=========================================#")
	fmt.Println("  Products")
	fmt.Println("-------------")
	for items := range product {
		fmt.Println("Product ID : ", product[items].ProdId)
		fmt.Println("Product Name : ", product[items].ProdName)
		fmt.Println("----------------------------")
	}

	var choice string
	var prodid int
	var price int

	for {
		fmt.Println("1 . Product Details")
		fmt.Println("If you want to view product details (yes/no)")
		_, _ = fmt.Scan(&choice)
		if choice == "yes" {
			fmt.Println("Enter Product ID : ")
			_, _ = fmt.Scan(&prodid)
			item := products.Products{}.GetproductByID(prodid)

			if item == -1 {
				fmt.Println("Product Not Found")
			} else {
				fmt.Printf("Product ID : %v\nProduct Name : %v\nPrice : %v\nQuantity : %v\n", product[item].ProdId, product[item].ProdName, product[item].Price, product[item].Quantity)
				fmt.Println("Order this Product (yes/no)")
				_, _ = fmt.Scan(&choice)

				if choice == "yes" {
					var quantity int
					fmt.Println("Enter Quantity : ")
					_, _ = fmt.Scan(&quantity)
					fmt.Printf("Enter Price (%v) : \n", products.ProductList[item].Price*quantity)
					_, _ = fmt.Scan(&price)
					wg.Add(2)
					go payment.Payment(products.ProductList[item].ProdId, quantity, price, &wg)
					wg.Wait()
				}
			}
		} else {
			break
		}
	}

}

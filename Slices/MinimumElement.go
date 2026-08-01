package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 8, 2, 4, 1, 5}

	minn := math.MaxInt64

	fmt.Println(minn)

	for _, num := range arr {
		if minn > num {

			minn = num
		}
	}

	fmt.Println(minn)
}

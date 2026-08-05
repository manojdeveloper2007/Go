package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func addFirst(arr []int, ch chan int) {
	sum := 0
	for i := 0; i < len(arr)/2; i++ {
		sum = sum + arr[i]
		fmt.Printf("Add First : %v and Sum : %v\n", arr[i], sum)
	}
	ch <- sum
	defer wg.Done()
}

func addSecond(arr []int, ch chan int) {
	sum := 0
	for i := len(arr) / 2; i < len(arr); i++ {
		sum += arr[i]
		fmt.Printf("Add Second : %v and Sum : %v\n", arr[i], sum)
	}
	ch <- sum
	defer wg.Done()
}

func main() {
	//creating channel
	ch := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg.Add(2)
	go addFirst(arr, ch)
	go addSecond(arr, ch)
	defer wg.Wait()
	total := <-ch + <-ch
	fmt.Println(total)
}

package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	arrayName := [5]string{"item1", "item2", "item3", "itemN"}
	fmt.Println(arrayName)
}

func bs(arrayName string, target int) int {
	return target
}

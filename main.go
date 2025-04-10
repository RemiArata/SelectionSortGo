package main

import (
	"SelectionSortGo/helpers"
	"SelectionSortGo/ss"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	someData := helpers.GenerateData()
	fmt.Println("The unsorted data is:", someData)

	sortedData := ss.SelectionSort(someData)
	fmt.Println("The sorted data is:", sortedData)
}

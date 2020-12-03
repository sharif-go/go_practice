package main

import "fmt"

func main() {

	arr := []string{"Nahid", "Faruk", "Tanzil", "Shagor", "Opu"}
	//Simple range
	for k, v := range arr {
		fmt.Println("Positon Number:", k, "is", v)
	}

	num := []string{"Nahid", "Faruk", "Towsif", "Tanzil", "Shagor", "Opu"}
	// Range with condition
	for i, k := range num {
		if k == "Shagor" {
			fmt.Println("Shagor's index number is:", i)
		}
	}
}

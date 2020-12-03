package main

import "fmt"

func test() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func main() {
	result := test()
	fmt.Println(result()) //output 1

	//affected by closures
	fmt.Println(result()) //output 2
	fmt.Println(result()) //output 3

	result2 := test()
	fmt.Println(result2()) //output 1

}

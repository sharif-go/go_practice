package main

import "fmt"

//Recursive function example 1
func test1(n int) {
	if n == 0 {
		fmt.Println(0)
	}
	fmt.Println(test1())
}

//Recursive function example 2
func recursiveFunction(num int) int {
	if num <= 1 {
		return num
	}
	return recursiveFunction(num-2) + recursiveFunction(num-1)

}

//Recursive function example 3
func recursiveFunction2(num int) int {
	if num == 0 {
		return 1
	}
	return num * recursiveFunction2(num-1)

}

func main() {
	fmt.Println(recursiveFunction(10))
	fmt.Println(recursiveFunction2(4))
	test1(10)
}

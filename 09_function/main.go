package main

import "fmt"

//simple function
func test() string {
	return ("This content show from test function")
}

//function with parameters
func test2(a, b int) int {
	return a + b
}

//recursive function
func test3(a int) int {
	if a <= 1 {
		return a
	} else {
		return test3(a-1) + test3(a-2)
	}
}

//Multiple return value
func test4() (int, string) {
	return 3, "test content"
}

func main() {

	//call function
	fmt.Println(test())
	fmt.Println(test2(23, 7))
	fmt.Println(test3(3))

	//Call  function with all agruments
	resultInt, resultString := test4()
	fmt.Println(resultString)
	fmt.Println(resultInt)

	//Call function with one argument
	_, arg := test4()
	fmt.Println(arg)
}

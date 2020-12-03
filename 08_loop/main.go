package main

import "fmt"

func main() {

	//loop with condition
	for i := 0; i < 10; i++ {
		fmt.Println("This line will print ten time")
	}

	//simple loop (note: Golang comes with just one loop that is called "for")
	for {
		fmt.Println("This line will print unlimited time")
	}
}

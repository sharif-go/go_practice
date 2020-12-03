package main

import "fmt"

func main() {
	marks := 32
	//simple switch
	switch {
	case marks > 32:
		fmt.Println("You are pass in the exam")
	case marks > 39:
		fmt.Println("Your GPA 2")
	case marks > 49:
		fmt.Println("Your GPA is up to 2")
	default:
		fmt.Println("You are failed in the exam")
	}
}

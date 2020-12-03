package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple if/else example 1
	num := 2
	if num > 5 {
		fmt.Println(num, "is greater then 5")
	} else {
		fmt.Println(num, "is less then 5")
	}

	//Simple if/else example 2
	currenttime := time.Now()

	if currenttime.Hour() > 12 {
		fmt.Println("Now is afternoon")
	} else {
		fmt.Println("Now is morning")
	}

	//if/elseif example

	marks := 32
	if marks >= 33 {
		fmt.Println("You are pass in the exam")
	} else if marks >= 40 {
		fmt.Println("Your result is GPA 2")
	} else if marks >= 50 {
		fmt.Println("Your result is 50 or more")
	} else {
		fmt.Println("You are failed in the exam")
	}
}

package main

import "fmt"

func main() {
	//Initialize map
	age := make(map[string]int)

	//implementation map
	age["Sharif"] = 20

	//print
	fmt.Println(age)

	//add another value
	age["Shamim"] = 16
	age["Rahul"] = 17
	age["Robin"] = 18

	fmt.Println(age)

	//access to particular value
	rahulAge := age["Rahul"]

	fmt.Println(rahulAge)

	//delete value from map
	delete(age, "Robin")

	fmt.Println(age)

	//access map by loop
	for k, v := range age {
		fmt.Println(k, ":", v)
	}

	//len() function with map
	getLength := len(age)

	fmt.Println("The map length is", getLength)

	//another way to implemention map
	age2 := map[string]int{"Bappy": 15, "Mosarrof": 23, "Rokib": 53}
	fmt.Println("Another map here:", age2)

	//value and weather the value is present
	_, is_value_present := age2["Bappy"]
	fmt.Println(is_value_present)
}

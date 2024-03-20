package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age
	var agePP **int = &agePointer

	fmt.Println(age)
	fmt.Println(&age)
	fmt.Println(agePointer)
	fmt.Println(&agePointer)
	fmt.Println(*agePointer)
	fmt.Println(agePP)
	fmt.Println(&agePP)
	fmt.Println(*agePP)
	fmt.Println(**agePP)

}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
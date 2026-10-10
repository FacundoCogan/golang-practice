package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age

	fmt.Println("Age: ", *agePointer)
	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {

	// return *age - 18

	*age = *age - 18

}

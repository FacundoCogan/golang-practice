package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {

	// 1
	hobbies := [3]string{"train", "work", "study"}

	fmt.Println("Original hobbies: ", hobbies)

	// 2
	fmt.Println("First element: ", hobbies[0]) // a

	fmt.Println("Second and third element: ", hobbies[1:]) // b

	// 3
	highlightedHobbies := hobbies[0:2] // b
	fmt.Println("First and second element: ", highlightedHobbies)

	// 4
	highlightedHobbies = highlightedHobbies[1:3]
	fmt.Println("Second and third element: ", highlightedHobbies)

	// 5
	courseGoals := []string{"Learn go!", "Learn all the basics"}
	fmt.Println("Course goals: ", courseGoals)

	// 6
	courseGoals[1] = "Lean all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println("New course goals: ", courseGoals)

	// 7
	products := []Product{{"A book", "1", 53.0}, {"A carpet", "2", 23.9}} // a
	fmt.Println("Products: ", products)

	products = append(products, Product{"A pen", "3", 34.9})

	fmt.Println("New products: ", products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

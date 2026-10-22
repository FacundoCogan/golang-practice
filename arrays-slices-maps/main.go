package main

import "fmt"

type floatMap map[string]float64 // alias

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // más eficiente para la memoria porque crea el array con capacidad de 5
	// deja los primeros dos espacios en blanco

	userNames[0] = "Julie" // así se asigna a uno de los espacios en blanco

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3) // lo mismo pero para maps

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	// fmt.Println(courseRatings)

	for index, value := range userNames { // si no quiero obtener index y valores, puedo ignorarlos y escribir for range userNames
		// ...
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	for key, value := range courseRatings {
		// ...
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}

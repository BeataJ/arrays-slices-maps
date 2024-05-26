package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "bob"
	// userNames[1] = "eve"

	userNames = append(userNames, "max")
	userNames = append(userNames, "beata")

	fmt.Println(userNames)

	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	courseRating["angular"] = 4.7

	courseRating.output()

	// fmt.Println(courseRating)

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRating {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}

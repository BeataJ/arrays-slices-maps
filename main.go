package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "bob"
	userNames[1] = "eve"

	userNames = append(userNames, "max")
	userNames = append(userNames, "beata")

	fmt.Println(userNames)

	courseRating := make(map[string]float64, 3)

	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	courseRating["angular"] = 4.7

	fmt.Println(courseRating)
}

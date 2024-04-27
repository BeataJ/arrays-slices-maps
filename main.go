package main

import "fmt"

func main() {
	userNames := []string{}

	userNames = append(userNames, "max")
	userNames = append(userNames, "beata")

	fmt.Println(userNames)
}

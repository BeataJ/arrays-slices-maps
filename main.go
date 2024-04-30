package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "bob"

	userNames = append(userNames, "max")
	userNames = append(userNames, "beata")

	fmt.Println(userNames)
}

package main

import "fmt"

func printMenu(menu map[string]float64) {
	for k, v := range menu {
		fmt.Println("In the menu the", k, "costs", v)
	}
}

func main() {

	menu := map[string]float64{
		"soup": 5.0,
		"pie":  10.50,
		"cake": 8.0,
	}

	fmt.Println("In the menu the soup costs", menu["soup"])

	printMenu(menu)

	menu["soup"] = 5.5

	printMenu(menu)
}

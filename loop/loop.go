package main

import "fmt"

func main() {
	names := []string{"luca", "andrea", "giulio"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println(name, " is at position ", i)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

package main

import "fmt"

func main() {

	names := []string{"luca", "andrea", "giulio"}

	for i, name := range names {
		if i == 1 {
			fmt.Println("Skipping value at position ", i)
			continue
		}

		fmt.Println(name, " is at position ", i)
	}

}

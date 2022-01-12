package main

import "fmt"

func main() {
	hi := "Hallo world from"
	name := "Luca"
	value := 0.3333333333

	fmt.Println(hi, name)

	fmt.Printf("Hallo world from %v \n", name)
	fmt.Printf("float %0.3f", value)
}

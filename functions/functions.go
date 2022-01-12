package main

import (
	"fmt"
	"math"
	"strings"
)

func cycleName(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func sayHallo(name string) {
	fmt.Println(name)
}

func calcArea(r float64) float64 {
	return math.Pi * r * r
}

func split(word string) (string, string) {

	split := strings.Split(word, " ")

	var r []string
	for _, w := range split {
		r = append(r, w)
	}

	return r[0], r[1]
}

func main() {
	cycleName([]string{"luca", "jimmy", "jhonny"}, sayHallo)
	fmt.Println(calcArea(10.5))

	// multiple returns

	w1, w2 := split("Hallo world")
	fmt.Println(w1)
	fmt.Println(w2)
}

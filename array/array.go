package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	hi := "Hallo world from"
	name := "Luca"
	fmt.Println(hi, name)

	fmt.Println(strings.Contains(hi, "Hallo"))

	var a [4]int
	a[0] = 1
	a[1] = 2
	a[2] = 0
	a[3] = 3

	fmt.Println("array initialized -> var a [4]int", a)

	ages := []int{2, 5, 6, 3}
	fmt.Println("array already with values -> ages := []int{2, 5, 6, 3}", ages)

	fmt.Println("slice of first 2 values -> ages[0:2]", ages[0:2])

	ages = append(ages, 4)
	fmt.Println("Append to array -> ", ages)

	sort.Ints(ages)

	fmt.Println("Sorted int array -> ", ages)

	names := []string{"luca", "andrea", "giulio"}

	fmt.Println("Array of strings", names)

	sort.Strings(names)

	fmt.Println("Sorted string array -> ", names)
}

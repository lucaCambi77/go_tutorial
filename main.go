package main

import (
	"fmt"
	"go_tutorial/struct"
)

func main() {

	myBill := _struct.CreateBill("Luca's bill")

	myBill.AddItem("soup", 5.00)
	myBill.AddItem("pie", 10.00)
	myBill.AddItem("cake", 8.00)

	fmt.Println(myBill.Format())
	myBill.AddItem("meat", 12.00)

	fmt.Println(myBill.Format())
}

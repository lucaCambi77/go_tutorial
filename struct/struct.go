package _struct

import "fmt"

func CreateBill(name string) Bill {

	bill := Bill{
		name:  name,
		items: map[string]float64{},
	}

	return bill
}

func (bill *Bill) Format() string {
	fs := "Bill Receipt \n"
	total := float64(0)

	for k, v := range bill.items {
		fs += fmt.Sprintf("%v ...$%v \n", k, v)
		total += v
	}

	fs += fmt.Sprintf("%v ...$%0.2f \n", "total:", total+bill.tip)

	return fs
}

func (bill *Bill) AddItem(name string, price float64) {
	bill.items[name] = price
}

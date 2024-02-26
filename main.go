package main

import (
	"fmt"
)

func main() {
	Prod3 := []Prod2{
		{"Kayak", 34.45},
		{"Bolo", 356.23},
		{"steart", 23.12},
	}

	NewProd2sSort(Prod3)
	fmt.Println("Sorted by price: ", Prod3)

	// UsersStructSort sort StrProd2 by name
	UsersStructSort(Prod3, func(p1, p2 Prod2) bool {
		return p1.name < p2.name
	})
	fmt.Println("Sorted by name: ", Prod3)

}

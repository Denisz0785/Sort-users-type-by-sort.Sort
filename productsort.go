package main

import "sort"

// create users type of structure
type Prod2 struct {
	name  string
	price float64
}
type NewProd2 []Prod2

// realize NewProd2 for method Less of interface
func (p NewProd2) Less(i, j int) bool {
	return p[i].price < p[j].price
}

// realize NewProd2 for method Len of interface
func (p NewProd2) Len() int {
	return len(p)
}

// realize NewProd2 for method Swap of interface
func (p NewProd2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// NewProd2Sort sorts []Prod2 by price
func NewProd2sSort(p []Prod2) {
	sort.Sort(NewProd2(p))

}

// create new type Compare
type Compare func(p1, p2 Prod2) bool

// create new struct
type StrProd2 struct {
	NewProd2
	Compare
}

// realize StrProd2 for method Less of interface
func (p StrProd2) Less(i, j int) bool {
	return p.Compare(p.NewProd2[i], p.NewProd2[j])

}

// UsersStructSort sorts StrProd2
func UsersStructSort(p NewProd2, f Compare) {
	sort.Sort(StrProd2{p, f})
}

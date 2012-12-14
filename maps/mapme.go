package main

import "fmt"

type structorx struct {
	x, y int
	z    string
}

func main() {
	// first just play with a simple map
	m1 := make(map[string]int)
	m1["test"] = 3
	m1["test2"] = 4
	fmt.Println(m1)
	// ok that was nice, so now a map to a struct
	m := make(map[string]*structorx)
	a := new(structorx)
	a.x, a.y, a.z = 1, 2, "three"
	b := new(structorx)
	b.x, b.y, b.z = 4, 5, "six"
	m["t1"] = a
	m["t2"] = b
	fmt.Println(m)
	fmt.Println(m["t1"].z)
}

package main

import "fmt"

func main() {
	// make 2 slices, one 5 characters long (5 cap) (filled with nil or 0 in int)
	// the second with no length but the capacity to grow to 7
	a := make([]int, 5)
	b := make([]int, 0, 7)
	fmt.Println("a:", a, "\nb:", b)
	fmt.Println("capacity (a) : length (a):", cap(a), ":", len(a))
	fmt.Println("capacity (b) : length (b):", cap(b), ":", len(b))
	// set the 3rd element in the slice to the # 3
	a[2] = 3
	// add 1 and 2 as two elements in the slice
	b = append(b, 1, 2)
	fmt.Println("a:", a, "\nb:", b)
	fmt.Println("capacity (a) : length (a):", cap(a), ":", len(a))
	fmt.Println("capacity (b) : length (b):", cap(b), ":", len(b))
	// here we are making the slice literally and assigning the vals
	c := make([]int, 0, 10)
	c = append(c, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// check the values
	fmt.Println("c val: ", c, "c cap:", cap(c), "c len:", len(c))
	var x int
	// loop through the slice popping off each one from the end
	for len(c) > 0 {
		x, c = c[len(c)-1], c[:len(c)-1]
		fmt.Println("pop'ed off val: ", x)
		fmt.Println("c val: ", c, "c cap:", cap(c), "c len:", len(c))
	}
	c = append(c, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// ok, reset the data in the slice, this time we are going to take it from the front
        for len(c) > 0 {
                x, c = c[0], c[1:]
                fmt.Println("pop'ed off val: ", x)
                fmt.Println("c val: ", c, "c cap:", cap(c), "c len:", len(c))
        }
}

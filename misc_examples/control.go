/*

I suppose I should get used to the proper layout, so here is my block comment which would contain cr info

*/

package main

import "fmt"

func main() {
	var x int  // declare x to be a variable of the type int
	x = 3      // set a value for x
	y := 1 + x // create variable y and set its value to 1 + x (which would make y an int)

	// a simple if
	if y > 0 {
		fmt.Println("yep")
	}
	fmt.Println(x)
	fmt.Println(y)

	// now an if/else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Println("no clue how that could be?")
	}

	// how about those bools?
	if true {
		fmt.Println("yes true is true")
	}

	/*
		options for comparision;
		== equal to
		!= not equal to
		< less than 
		<= less than or equal to 
		> greater than
		>= greater than or equal to
		&& and
		|| or
	*/

	// all good but nothing shocking, how about the accept and init bit?
	if z := x + y; z > x {
		fmt.Println("better be right")
	} else {
		fmt.Println("really?  how?")
	}

	// for loops
	fmt.Println("now try the for loop")
	n := 0
	for i := 0; i < 10; i++ {
		n += i
		fmt.Printf("VALS: %d - %d\n", i, n)
	}

	// switch  (oh my oh my a switch statement, imagine that)
	a, b := 3, 4
	fmt.Printf("a=%d, b=%d\n", a, b)
	switch {
	case a < 2:
		fmt.Println("less than 2?")
	case a == 3:
		fmt.Println("better be")
		// note this could be a return, etc
	}
	fmt.Println("all done")

}

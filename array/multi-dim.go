package main

import (
	"fmt"
	
)

var n int = 1

func main() {
	m := [3][4]string {}
	m[0][0] = "aba"
	m[0][1] = "bcd"
	fmt.Println(m)

	n := [4][4]int{}
	n[0][0]=1
	n[0][1]=10
	n[1][0]=4
	n[1][1]=6
	n[2][0]=7
	n[2][1]=3
	fmt.Println(n)
	mat := [6][6]int{}
	fmt.Println(mat)
}
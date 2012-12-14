package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	a, err := regexp.Compile(`[a-z]`)
	d, err := regexp.Compile(`[1-9]`)
	if err != nil {
		fmt.Println("BAD REGEX")
		os.Exit(-1)
	}

	if a.MatchString(os.Args[1]) {
		fmt.Printf("a: matched on %s \n", os.Args[1])
	} else if d.MatchString(os.Args[1]) {
		fmt.Printf("d: matched on %s \n", os.Args[1])
	} else {
		fmt.Println("no match")
	}
}

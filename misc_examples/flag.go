package main

import (
	"flag"
	"fmt"
	"os"
)

var age *int = flag.Int("age", -1, "you must enter your age")
var name *string = flag.String("name", "", "you must enter your name")

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	// PrintDefaults() may not be exactly what we want, but it could be
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	fmt.Printf("testing the flag package\n")
	flag.Parse()
	if *age == -1 || *name == "" {
		usage()
	}
	fmt.Printf("Age: %d\nName:%s\n", *age, *name)
}

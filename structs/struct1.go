package main

import "fmt"

type structorx struct {
        x, y int
        z    string
}

func main() {
        a := new(structorx)
        fmt.Println(a)
        a.x, a.y, a.z = 1, 2, "three"
        fmt.Println(a)
}

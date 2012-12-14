package main

import "fmt"

type structorx struct {
        x, y int
        z    string
}

func main() {
        a := new(structorx)
        a.x, a.y, a.z = 1, 2, "three"
        b := new(structorx)
        b.x, b.y, b.z = 4, 3, "two"
        fmt.Println(a)
        fmt.Println(b)
	c := make([]*structorx, 1)
        c[0] = a
        c = append(c, b)
        fmt.Println(c)
        fmt.Println(c[0].x, c[0].y, c[0].z)
        fmt.Println(c[1].x, c[1].y, c[1].z)
}
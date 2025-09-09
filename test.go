package main

import (
	"fmt"
)

func main() {
	var (
		a int = 9
		b     = "b"
		c bool
		e string
	)
	c = true
	e = "not ok"
	d, f := 9, "false"

	fmt.Println(a, b, c, d, e, f)

}

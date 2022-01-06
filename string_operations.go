package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "helloworld hellouniverse"
	p := "   " + a + "   "
	// trim a string
	b := strings.Trim(a, "he")
	c := strings.TrimLeft(a, "hel")
	d := strings.TrimRight(a, "rse")
	e := strings.TrimSuffix(a, "verse")
	f := strings.TrimPrefix(a, "hel")
	s := strings.TrimSpace(p)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(s)

}

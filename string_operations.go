package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hell oworld hello universe"
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
	// split a string same as python
	// split
	g := strings.Split(a, " ") // unlike python Split function in golang doen't have default arguments as " " so we have to specify it manually...
	fmt.Println(g)
	// split after
	h := strings.SplitAfter(a, " ") // so result have space included in it
	fmt.Println(h)
	// split after n
	i := strings.SplitAfterN(a, " ", 2) // so result have space included in it
	fmt.Println(i)

	// comapring two strings in lexciographical manner
	str1 := "Sonali"
	str2 := "Mukul"
	str3 := "Hi"
	str4 := "Sonali"
	// Normal check to check if they are equal
	c1 := str1 == str2
	c2 := str2 == str3
	c3 := str1 == str4
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	// Now lets check which one is  dominant over other
	c4 := str1 > str2
	fmt.Println(c4, "Sonali is dominant over Mukul")
	// we can also compare using strings.Compare function
	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Compare(str2, str1))
	fmt.Println(strings.Compare(str1, str4))
	// As we can see using compare function we get 0 if it holds equality, 1 if 1st argument is greater than second argument and -1 if second argument is greater than first argument.
}

package main

import (
	"fmt"
)

func main() {
	//Declare variables that're set to 0 values.
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var a string \t %T [%v]\n", b, b)
	fmt.Printf("var a float64 \t %T [%v]\n", c, c)
	fmt.Printf("var a boolean \t %T [%v]\n", d, d)

	//Declare variables & initialize
	//using short variable delaration operator
	aa := 10
	bb := "hare kRSNa"
	cc := 1.2345
	dd := false
	fmt.Print("%T [%v]\n%T [%v]\n%T [%v]\n %T [%v]", aa, aa, bb, bb, cc, cc, dd, dd)

}

package main

import (
	"math"
)

func Return3() int {
	return 3
}

func Add(x, y int) {
	return x + y
}
func Sub(x, y int) {
	return x - y
}

func Div(x, y int) {
	return x / y
}

func Mod(x, y int) {
	return x % y
}

func Mul(x, y int) {
	return x * y
}

func Pow(x, y int) {
	return math.pow(x, y)
}

func Sqrt(x int) {
	return math.pow(x, (1 / 2))
}

func Curt(x int) {
	return math.pow(x, (1 / 3))
}

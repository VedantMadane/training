package main

import "testing"

const succeed = "\033[32m\u2713"
const failure = "\033[31m\u2717"
const normal = "\033[0m"

type binfunc func(float32, float32) float32

type BinaryOp struct {
	x   float32
	y   float32
	ans float32
}

func runTestCases(t *testing.T, testcases []BinaryOp, function binfunc, title string) {
	t.Log(title)
	{
		for i, tt := range testcases {
			t.Logf("\tTest: %d | func(%f, %f) == %f", i, tt.x, tt.y, tt.ans)
			ans := function(tt.x, tt.y)
			if ans != tt.ans {
				t.Fatalf("\t%s Should return correct ans \n\texpected %f \n\treceived %f%s", failure, tt.ans, ans, normal)
				return
			}
			t.Logf("\t%s Passed %s", succeed, normal)
		}
	}
}

func TestReturn3(t *testing.T) {
	ret := Return3()
	if ret == 3 {
		t.Logf("\n%s Returned correct ans %s", succeed, normal)
	} else {
		t.Fatalf("\n%s Returned wrong ans %s", failure, normal)
		return
	}
}

func TestAddition(t *testing.T) {
	testcases := []BinaryOp{
		{x: 3, y: 4, ans: 7},
		{x: 4, y: 3, ans: 7},
		{x: -2, y: -2, ans: -4},
		{x: 0, y: 4, ans: 4},
		{x: 12, y: 104, ans: 116},
	}
	runTestCases(t, testcases, Addition, "Should be able to add numbers")
}

func TestSubstraction(t *testing.T) {
	testcases := []BinaryOp{
		{x: 3, y: 4, ans: -1},
		{x: 4, y: 3, ans: 1},
		{x: -2, y: -2, ans: 0},
		{x: 0, y: 4, ans: -4},
		{x: 12, y: 104, ans: -92},
	}
	runTestCases(t, testcases, Substraction, "Should be able to subtract numbers")
}

func TestDivision(t *testing.T) {
	testcases := []BinaryOp{
		{x: 3, y: 4, ans: 0.75},
		{x: 4, y: 3, ans: 4.0 / 3.0},
		{x: -2, y: -2, ans: 1},
		{x: 0, y: 4, ans: 0},
		{x: 120, y: 12, ans: 10},
	}
	runTestCases(t, testcases, Division, "Should be able to divide numbers")
}

func TestMultiplication(t *testing.T) {
	testcases := []BinaryOp{
		{x: 3, y: 4, ans: 12},
		{x: 4, y: 3, ans: 12},
		{x: -2, y: -2, ans: 4},
		{x: 0, y: 4, ans: 0},
		{x: 12, y: 12, ans: 144},
	}
	runTestCases(t, testcases, Multiplication, "Should be able to multiply numbers")
}

package helpers

import (
	"fmt"
	"testing"
)

type CaseMath struct {
	A      int
	B      int
	Return int
}

// region Min

var testCasesMin = []CaseMath{
	{A: 0, B: 0, Return: 0},
	{A: 42, B: 0, Return: 0},
	{A: 0, B: 42, Return: 0},
	{A: -42, B: 0, Return: -42},
	{A: 0, B: -42, Return: -42},
	{A: 42, B: -42, Return: -42},
	{A: -42, B: 42, Return: -42},
}

func TestMin(t *testing.T) {
	for _, testCase := range testCasesMin {
		if value := Min(testCase.A, testCase.B); value != testCase.Return {
			t.Error("A:", testCase.A, "B:", testCase.B, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMin(b *testing.B) {
	var testCase CaseMath
	for i := 0; i < b.N; i++ {
		testCase = testCasesMin[i%len(testCasesMin)]
		Min(testCase.A, testCase.B)
	}
}

func ExampleMin() {
	fmt.Println(Min(-42, 42))
	// Output: -42
}

// endregion

// region Max

var testCasesMax = []CaseMath{
	{A: 0, B: 0, Return: 0},
	{A: 42, B: 0, Return: 42},
	{A: 0, B: 42, Return: 42},
	{A: -42, B: 0, Return: 0},
	{A: 0, B: -42, Return: 0},
	{A: 42, B: -42, Return: 42},
	{A: -42, B: 42, Return: 42},
}

func TestMax(t *testing.T) {
	for _, testCase := range testCasesMax {
		if value := Max(testCase.A, testCase.B); value != testCase.Return {
			t.Error("A:", testCase.A, "B:", testCase.B, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMax(b *testing.B) {
	var testCase CaseMath
	for i := 0; i < b.N; i++ {
		testCase = testCasesMax[i%len(testCasesMax)]
		Max(testCase.A, testCase.B)
	}
}

func ExampleMax() {
	fmt.Println(Max(-42, 42))
	// Output: 42
}

// endregion

package main

import "fmt"

type CalculatorNoDIP struct{}

func (c *CalculatorNoDIP) Add(a, b int) int {
	return a + b
}

func (c *CalculatorNoDIP) Subtract(a, b int) int {
	return a - b
}

func (c *CalculatorNoDIP) Multiply(a, b int) int {
	return a * b
}

func (c *CalculatorNoDIP) Divide(a, b int) int {
	return a / b
}

func main() {
	// no dip
	calc := CalculatorNoDIP{}
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(1, 2))
	fmt.Println(calc.Multiply(1, 2))
	fmt.Println(calc.Divide(1, 2))
}

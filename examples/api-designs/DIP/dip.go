package main

import "fmt"

type Calculator interface {
	Calculate(a, b float64) float64
}

type Adder struct{}

func (a Adder) Calculate(x, y float64) float64 {
	return x + y
}

type Subtract struct{}

func (s Subtract) Calculate(x, y float64) float64 {
	return x - y
}

type Multiplier struct{}

func (m Multiplier) Calculate(x, y float64) float64 {
	return x * y
}

type Divider struct{}

func (d Divider) Calculate(x, y float64) float64 {
	return x / y
}

type CalculatorDIP struct {
	calculator Calculator
}

func NewCalculatorDIP(calculator Calculator) *CalculatorDIP {
	return &CalculatorDIP{calculator: calculator}
}

func (c *CalculatorDIP) Operate(a, b float64) float64 {
	return c.calculator.Calculate(a, b)
}

func main() {
	adder := Adder{}
	subtract := Subtract{}
	multiplier := Multiplier{}
	divider := Divider{}

	calculator := NewCalculatorDIP(adder)
	fmt.Println(calculator.Operate(1, 2))

	calculator.calculator = subtract
	fmt.Println(calculator.Operate(1, 2))

	calculator.calculator = multiplier
	fmt.Println(calculator.Operate(1, 2))

	calculator.calculator = divider
	fmt.Println(calculator.Operate(1, 2))
}
package calculator

import (
	"fmt"
	"math"
)

type Calculator struct {
	stacks    [4]float64       // x, y, a, b
	vars      map[rune]float64 // A-Z
	angleMode string           // "RAD" or "DEG"
}

func NewCalculator() *Calculator {
	return &Calculator{
		stacks:    [4]float64{0, 0, 0, 0},
		vars:      make(map[rune]float64),
		angleMode: "RAD",
	}
}

func (c *Calculator) push(val float64) {
	c.stacks[3] = c.stacks[2]
	c.stacks[2] = c.stacks[1]
	c.stacks[1] = c.stacks[0]
	c.stacks[0] = val
}

func (c *Calculator) pop() float64 {
	val := c.stacks[0]
	c.stacks[0] = c.stacks[1]
	c.stacks[1] = c.stacks[2]
	c.stacks[2] = c.stacks[3]
	c.stacks[3] = 0
	return val
}

func (c *Calculator) peek() float64 {
	return c.stacks[0]
}

func (c *Calculator) operate(op string) {
	switch op {
	case "+":
		b := c.pop()
		a := c.pop()
		c.push(a + b)
	case "-":
		b := c.pop()
		a := c.pop()
		c.push(a - b)
	case "*":
		b := c.pop()
		a := c.pop()
		c.push(a * b)
	case "/":
		b := c.pop()
		a := c.pop()
		c.push(a / b)
	case "sqrt":
		a := c.pop()
		c.push(math.Sqrt(a))
	case "sq":
		a := c.pop()
		c.push(a * a)
	case "exp":
		a := c.pop()
		c.push(math.Exp(a))
	case "pow":
		b := c.pop()
		a := c.pop()
		c.push(math.Pow(a, b))
	case "sin":
		a := c.pop()
		if c.angleMode == "DEG" {
			a = a * math.Pi / 180
		}
		c.push(math.Sin(a))
	case "cos":
		a := c.pop()
		if c.angleMode == "DEG" {
			a = a * math.Pi / 180
		}
		c.push(math.Cos(a))
	case "tan":
		a := c.pop()
		if c.angleMode == "DEG" {
			a = a * math.Pi / 180
		}
		c.push(math.Tan(a))
	case "RAD":
		c.angleMode = "RAD"
		fmt.Println("Angle mode set to radians.")
	case "DEG":
		c.angleMode = "DEG"
		fmt.Println("Angle mode set to degrees.")
	case "clear":
		for i := range c.stacks {
			c.stacks[i] = 0
		}
		fmt.Println("Stack cleared.")
	}
}

func (c *Calculator) store(varName rune) {
	if varName >= 'A' && varName <= 'Z' {
		c.vars[varName] = c.peek()
		fmt.Printf("Stored %.8g in %c\n", c.peek(), varName)
	}
}

func (c *Calculator) recall(varName rune) {
	if val, ok := c.vars[varName]; ok {
		c.push(val)
		fmt.Printf("Recalled %.8g from %c\n", val, varName)
	} else {
		fmt.Printf("No value stored in %c\n", varName)
	}
}

func (c *Calculator) PrintStacks() {
	fmt.Printf("x: %.8g\ty: %.8g\ta: %.8g\tb: %.8g\n", c.stacks[0], c.stacks[1], c.stacks[2], c.stacks[3])
}

func (c *Calculator) StackValues() [4]float64 {
	return c.stacks
}

func (c *Calculator) Push(val float64)    { c.push(val) }
func (c *Calculator) Pop() float64        { return c.pop() }
func (c *Calculator) Peek() float64       { return c.peek() }
func (c *Calculator) Operate(op string)   { c.operate(op) }
func (c *Calculator) Store(varName rune)  { c.store(varName) }
func (c *Calculator) Recall(varName rune) { c.recall(varName) }

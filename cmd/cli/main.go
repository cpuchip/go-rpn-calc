package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cpuchip/go-rpn-calc/internal/calculator"
)

func main() {
	calc := calculator.NewCalculator()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("RPN Calculator. Enter numbers, operations, or commands. Type 'exit' to quit.")
	for {
		calc.PrintStacks()
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		fields := strings.Fields(line)
		for _, field := range fields {
			if v, err := strconv.ParseFloat(field, 64); err == nil {
				calc.Push(v)
			} else if len(field) == 4 && strings.ToUpper(field[:3]) == "STO" {
				calc.Store(rune(strings.ToUpper(field[3:])[0]))
			} else if len(field) == 4 && strings.ToUpper(field[:3]) == "REC" {
				calc.Recall(rune(strings.ToUpper(field[3:])[0]))
			} else if field == "STO" || field == "REC" {
				fmt.Println("Usage: STOA or RECA (A-Z)")
			} else {
				calc.Operate(strings.ToLower(field))
			}
		}
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/cpuchip/go-rpn-calc/internal/calculator"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("RPN Calculator (GUI)")

	calc := calculator.NewCalculator()
	stackLabel := widget.NewLabel(stackString(calc.StackValues()))
	outputLabel := widget.NewLabel("")
	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Enter RPN expression or command (e.g. 2 3 +)")

	inputEntry.OnSubmitted = func(text string) {
		if strings.TrimSpace(text) == "exit" {
			myApp.Quit()
			return
		}
		fields := strings.Fields(text)
		for _, field := range fields {
			if v, err := strconv.ParseFloat(field, 64); err == nil {
				calc.Push(v)
			} else if len(field) == 4 && strings.ToUpper(field[:3]) == "STO" {
				calc.Store(rune(strings.ToUpper(field[3:])[0]))
			} else if len(field) == 4 && strings.ToUpper(field[:3]) == "REC" {
				calc.Recall(rune(strings.ToUpper(field[3:])[0]))
			} else if field == "STO" || field == "REC" {
				outputLabel.SetText("Usage: STOA or RECA (A-Z)")
			} else {
				calc.Operate(strings.ToLower(field))
			}
		}
		stackLabel.SetText(stackString(calc.StackValues()))
		inputEntry.SetText("")
	}

	content := container.NewVBox(
		widget.NewLabel("RPN Calculator (GUI)"),
		stackLabel,
		outputLabel,
		inputEntry,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(480, 180))
	w.ShowAndRun()
}

func stackString(stacks [4]float64) string {
	return fmt.Sprintf("x: %.8g\ty: %.8g\ta: %.8g\tb: %.8g", stacks[0], stacks[1], stacks[2], stacks[3])
}

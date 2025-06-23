package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cpuchip/go-rpn-calc/internal/calculator"
)

type model struct {
	calc     *calculator.Calculator
	input    string
	output   string
	quitting bool
}

func initialModel() model {
	return model{
		calc: calculator.NewCalculator(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit
		case tea.KeyEnter:
			cmd := m.handleInput()
			m.input = ""
			return m, cmd
		case tea.KeyBackspace, tea.KeyDelete:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			return m, nil
		default:
			m.input += msg.String()
			return m, nil
		}
	}
	return m, nil
}

func (m *model) handleInput() tea.Cmd {
	if strings.TrimSpace(m.input) == "exit" {
		m.quitting = true
		return tea.Quit
	}
	fields := strings.Fields(m.input)
	for _, field := range fields {
		if v, err := strconv.ParseFloat(field, 64); err == nil {
			m.calc.Push(v)
		} else if len(field) == 4 && strings.ToUpper(field[:3]) == "STO" {
			m.calc.Store(rune(strings.ToUpper(field[3:])[0]))
		} else if len(field) == 4 && strings.ToUpper(field[:3]) == "REC" {
			m.calc.Recall(rune(strings.ToUpper(field[3:])[0]))
		} else if field == "STO" || field == "REC" {
			m.output = "Usage: STOA or RECA (A-Z)"
		} else {
			m.calc.Operate(strings.ToLower(field))
		}
	}
	return nil
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}
	stacks := m.calc.StackValues()
	return fmt.Sprintf(
		"RPN Calculator (TUI)\nx: %.8g\ty: %.8g\ta: %.8g\tb: %.8g\n%s\n> %s",
		stacks[0], stacks[1], stacks[2], stacks[3],
		m.output,
		m.input,
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}

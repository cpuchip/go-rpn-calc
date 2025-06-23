package calculator

import (
	"math"
	"testing"
)

func TestPushPop(t *testing.T) {
	c := NewCalculator()
	c.Push(42)
	if got := c.Pop(); got != 42 {
		t.Errorf("Pop() = %v, want 42", got)
	}
}

func TestAdd(t *testing.T) {
	c := NewCalculator()
	c.Push(2)
	c.Push(3)
	c.Operate("+")
	if got := c.Peek(); got != 5 {
		t.Errorf("2 + 3 = %v, want 5", got)
	}
}

func TestSub(t *testing.T) {
	c := NewCalculator()
	c.Push(5)
	c.Push(3)
	c.Operate("-")
	if got := c.Peek(); got != 2 {
		t.Errorf("5 - 3 = %v, want 2", got)
	}
}

func TestMul(t *testing.T) {
	c := NewCalculator()
	c.Push(4)
	c.Push(3)
	c.Operate("*")
	if got := c.Peek(); got != 12 {
		t.Errorf("4 * 3 = %v, want 12", got)
	}
}

func TestDiv(t *testing.T) {
	c := NewCalculator()
	c.Push(12)
	c.Push(3)
	c.Operate("/")
	if got := c.Peek(); got != 4 {
		t.Errorf("12 / 3 = %v, want 4", got)
	}
}

func TestSqrt(t *testing.T) {
	c := NewCalculator()
	c.Push(9)
	c.Operate("sqrt")
	if got := c.Peek(); got != 3 {
		t.Errorf("sqrt(9) = %v, want 3", got)
	}
}

func TestSinDeg(t *testing.T) {
	c := NewCalculator()
	c.Operate("DEG")
	c.Push(90)
	c.Operate("sin")
	if math.Abs(c.Peek()-1) > 1e-9 {
		t.Errorf("sin(90 DEG) = %v, want 1", c.Peek())
	}
}

func TestStoreRecall(t *testing.T) {
	c := NewCalculator()
	c.Push(7)
	c.Store('A')
	c.Push(0)
	c.Recall('A')
	if got := c.Peek(); got != 7 {
		t.Errorf("Recall('A') = %v, want 7", got)
	}
}

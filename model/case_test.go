package model

import (
	"fmt"
	"testing"
)

func TestQuestion(t *testing.T) {
	q := NewQuestion("What is your temperature?")
	fmt.Println(q)
	f := NewFact(q)
	fmt.Println(f)
}

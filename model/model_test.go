package model

import (
	"testing"
	"fmt"
)

func assertBool(t *testing.T, gold bool, test bool) {
	if gold != test {
		t.Fail()
	}
}

func assertObject(t *testing.T, gold interface{}, test interface{}) {
	if gold != test {
		fmt.Printf("Assert failed: %v != %v\n", gold, test)
		t.Fail()
	}
}

func assertInt(t *testing.T, gold int, test int) {
	if gold != test {
		t.Fail()
	}
}
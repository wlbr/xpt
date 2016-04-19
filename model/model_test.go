package model

import "testing"

func assertBool(t *testing.T, gold bool, test bool) {
	if gold != test {
		t.Fail()
	}
}

func assertObject(t *testing.T, gold interface{}, test interface{}) {
	if gold != test {
		t.Fail()
	}
}

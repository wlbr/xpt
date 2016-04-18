package model

import "testing"

func TestSimpleRule(t *testing.T) {

	a := And(Equals(Variable(7), Variable(6)), Equals(Variable(5), Variable(5)))
}

package model

import "testing"

func TestSimpleBoolInterpreterRule(t *testing.T) {
	a1 := And(Equals(Variable(true), Variable(true)), Equals(Variable(true), Variable(false)))
	a2 := And(Equals(Variable(true), Variable(true)), Variable(true))
	a3 := And(Equals(Variable(true), Variable(true)), Variable(false))

	o1 := Or(Equals(Variable(true), Variable(true)), Equals(Variable(true), Variable(false)))
	o2 := Or(Equals(Variable(true), Variable(true)), Variable(false))
	o3 := Or(Equals(Variable(false), Variable(true)), Variable(false))

	x1 := Xor(Equals(Variable(true), Variable(true)), Equals(Variable(true), Variable(true)))
	x2 := Xor(Equals(Variable(false), Variable(true)), Equals(Variable(true), Variable(true)))
	x3 := Xor(Equals(Variable(false), Variable(false)), Equals(Variable(false), Variable(false)))

	n1 := Not(Variable(true))
	n2 := Not(Variable(false))
	n3 := Not(x3)

	v := make(map[string]Expression)

	assertBool(t, false, a1.Interprete(v))
	assertBool(t, true, a2.Interprete(v))
	assertBool(t, false, a3.Interprete(v))

	assertBool(t, true, o1.Interprete(v))
	assertBool(t, true, o2.Interprete(v))
	assertBool(t, false, o3.Interprete(v))

	assertBool(t, false, x1.Interprete(v))
	assertBool(t, true, x2.Interprete(v))
	assertBool(t, false, x3.Interprete(v))

	assertBool(t, false, n1.Interprete(v))
	assertBool(t, true, n2.Interprete(v))
	assertBool(t, true, n3.Interprete(v))
}

package model

/* we deviate from the Constructor convention (New...) by intention, to enable concise
   expressions like
	a := And(Equals(Variable(7), Variable(6)), Equals(Variable(5), Variable(5)))
*/
type Expression interface {
	Interprete(map[string]Expression) bool
	//Interprete(Map<String,Expression> variables) bool
}

type VariableExpression struct {
	Value bool
}

func Variable(value bool) *VariableExpression {
	return &VariableExpression{Value: value}
}

func (s *VariableExpression) Interprete(variables map[string]Expression) bool {
	return s.Value
}

type EqualsExpression struct {
	leftOperand  Expression
	rightOperand Expression
}

func Equals(leftOperand Expression, rightOperand Expression) *EqualsExpression {
	return &EqualsExpression{leftOperand: leftOperand, rightOperand: rightOperand}
}

func (s *EqualsExpression) Interprete(variables map[string]Expression) bool {
	return s.leftOperand.Interprete(variables) == s.rightOperand.Interprete(variables)
}

type OrExpression struct {
	leftOperand  Expression
	rightOperand Expression
}

func Or(leftOperand Expression, rightOperand Expression) *OrExpression {
	return &OrExpression{leftOperand: leftOperand, rightOperand: rightOperand}
}

func (s *OrExpression) Interprete(variables map[string]Expression) bool {
	return s.leftOperand.Interprete(variables) || s.rightOperand.Interprete(variables)
}

type AndExpression struct {
	leftOperand  Expression
	rightOperand Expression
}

func And(leftOperand Expression, rightOperand Expression) *AndExpression {
	return &AndExpression{leftOperand: leftOperand, rightOperand: rightOperand}
}

func (s *AndExpression) Interprete(variables map[string]Expression) bool {
	return s.leftOperand.Interprete(variables) && s.rightOperand.Interprete(variables)
}

type XorExpression struct {
	leftOperand  Expression
	rightOperand Expression
}

func Xor(leftOperand Expression, rightOperand Expression) *XorExpression {
	return &XorExpression{leftOperand: leftOperand, rightOperand: rightOperand}
}

func (s *XorExpression) Interprete(variables map[string]Expression) bool {
	return s.leftOperand.Interprete(variables) != s.rightOperand.Interprete(variables)
}

type NotExpression struct {
	operand Expression
}

func Not(operand Expression) *NotExpression {
	return &NotExpression{operand: operand}
}

func (s *NotExpression) Interprete(variables map[string]Expression) bool {
	return !s.operand.Interprete(variables)
}

func (s Fact) Interprete(variables map[string]Expression) bool {
	return true
}

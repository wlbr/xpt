package model

type Expression interface {
	Interprete(map[string]Expression) bool
	//Interprete(Map<String,Expression> variables) bool
}

type VariableExpression struct {
	Value int
}

func Variable(value int) *VariableExpression {
	return &VariableExpression{Value: value}
}

func (s *VariableExpression) Interprete(variables map[string]Expression) int {
	return s.Value
}

func (s *VariableExpression) Interprete(variables map[string]Expression) bool {
	return true
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

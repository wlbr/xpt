package model

import "fmt"

type Rule struct {
	AbstractKnowledgebaseObject
	condition Expression
	action    Action
}

func (rule *Rule) String() string {
	return fmt.Sprintf("Rule: %s", rule.id)
}

type Action interface {
	Act()
}

type SuspectThesisAction struct {
	thesis   Thesis
	analysis Analysis
}

func SuspectThesis(t Thesis, a Analysis) *SuspectThesisAction {
	return &SuspectThesisAction{thesis: t, analysis: a}
}

func (s *SuspectThesisAction) Act() {

}

//--------------------

func (s Fact) Interprete(variables map[string]Expression) bool {
	return true
}

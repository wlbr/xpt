package model

import (
	"fmt"
	"xpt/tools"
	// "math/big"
)

type Knowledgebase struct {
	QuestionGroups []KnowledgebaseObject
	Theses         []Thesis
}

type KnowledgebaseObject interface {
	String() string
}

type AbstractKnowledgebaseObject struct {
	id string
}

func GenerateID(i interface{}) string {
	var key string
	switch i.(type) {
	default:
		key = "K"
	case *Question:
		key = "Q"
	case *Thesis:
		key = "T"
	case *Fact:
		key = "F"
	case *Analysis:
		key = "A"
	case *Rule:
		key = "R"
	}
	id := tools.GenSID(key)
	return id

}

type Question struct {
	AbstractKnowledgebaseObject
	text string
}

func NewQuestion(text string) *Question {
	q := &Question{text: text}
	q.id = GenerateID(q)
	return q
}

func (question *Question) String() string {
	return fmt.Sprintf("Question %s: %s", question.id, question.text)
}

type Option struct {
	text string
}

func NewOption(text string) *Option {
	return &Option{text: text}
}

type OcQuestion struct {
	Question
	opts []*Option
}

func NewOcQuestion(text string) *OcQuestion {
	oc := &OcQuestion{}
	q := NewQuestion(text)
	oc.Question = *q
	return oc
}

func (oc *OcQuestion) String() string {
	return fmt.Sprintf("OcQuestion %s: %s", oc.id, oc.text)
}

type McQuestion struct {
	Question
	opts1 []*Option
}

func NewMcQuestion(text string) *McQuestion {
	mc := &McQuestion{}
	q := NewQuestion(text)
	mc.Question = *q
	return mc
}

func (mc *McQuestion) String() string {
	return fmt.Sprintf("McQuestion %s: %s", mc.id, mc.text)
}

type TextQuestion struct {
	Question
	//answer string // TODO answer is a -->>fact
}

func NewTextQuestion(text string) *TextQuestion {
	tq := &TextQuestion{}
	q := NewQuestion(text)
	tq.Question = *q
	return tq
}

func (tq *TextQuestion) String() string {
	return fmt.Sprintf("TextQuestion %s: %s", tq.id, tq.text)
}

type NumQuestion struct {
	Question
	//answer int // TODO answer is a -->>fact
}

func NewNumQuestion(text string) *NumQuestion {
	nq := &NumQuestion{}
	q := NewQuestion(text)
	nq.Question = *q
	return nq
}

func (nq *NumQuestion) String() string {
	return fmt.Sprintf("NumQuestion %s: %s", nq.id, nq.text)
}

type Thesis struct {
	AbstractKnowledgebaseObject
	text string
}

func NewTheses(text string) *Thesis {
	t := &Thesis{text: text}
	t.id = GenerateID(t)
	return t
}

func (theses *Thesis) String() string {
	return fmt.Sprintf("Theses %s: %s", theses.id, theses.text)
}

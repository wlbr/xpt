package model

import (
	"fmt"
	"xpt/tools"
	// "math/big"
)

type Knowledgebase struct {
	QuestionGroups map[string]QuestionGroup
	allQuestions   map[string]Question
	allTheses      map[string]Thesis
	ThesisGroups   map[string]ThesisGroup
}

func NewKnowledgebase() *Knowledgebase {
	k := &Knowledgebase{}
	k.QuestionGroups = make(map[string]QuestionGroup)
	k.allQuestions = make(map[string]Question)
	k.allTheses = make(map[string]Thesis)
	k.ThesisGroups = make(map[string]ThesisGroup)

	return k
}

func (kb Knowledgebase) AddQuestionGroup(qg QuestionGroup) QuestionGroup {
	qt, exists := kb.QuestionGroups[qg.id]
	if exists == false {
		kb.QuestionGroups[qg.id] = qg
		qt = qg
	}
	return qt
}

func (kb Knowledgebase) AddThesisGroup(tg ThesisGroup) ThesisGroup {
	tt, exists := kb.ThesisGroups[tg.id]
	if exists == false {
		kb.ThesisGroups[tg.id] = tg
		tt = tg
	}
	return tt
}

func (kb Knowledgebase) AddThesis(tg ThesisGroup, t Thesis) {
	ntg := kb.AddThesisGroup(tg)
	ntg.AddThesis(t)

}

type KnowledgebaseObject interface {
	String() string
	Id() string
}

type AbstractKnowledgebaseObject struct {
	id string
}

func (ako *AbstractKnowledgebaseObject) Id() string {
	return ako.id
}

func GenerateID(i interface{}) string {
	var key string
	switch i.(type) {
	default:
		key = "K"
	case *abstractQuestion:
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

type QuestionGroup struct {
	AbstractKnowledgebaseObject
	text string
}

func NewQuestionGroup(text string) *QuestionGroup {
	qg := &QuestionGroup{text: text}
	qg.id = GenerateID(qg)
	return qg
}

func (qg *QuestionGroup) AddQuestion() {

}

type Question interface {
	String() string
	Id() string
}

type abstractQuestion struct {
	AbstractKnowledgebaseObject
	text string
}

func NewQuestion(text string) *abstractQuestion {
	q := &abstractQuestion{text: text}
	q.id = GenerateID(q)
	return q
}

func (question *abstractQuestion) String() string {
	return fmt.Sprintf("Question %s: %s", question.id, question.text)
}

type Option struct {
	text string
}

func NewOption(text string) *Option {
	return &Option{text: text}
}

type OcQuestion struct {
	abstractQuestion
	opts []*Option
}

func NewOcQuestion(text string) *OcQuestion {
	oc := &OcQuestion{}
	q := NewQuestion(text)
	oc.abstractQuestion = *q
	return oc
}

func (oc *OcQuestion) String() string {
	return fmt.Sprintf("OcQuestion %s: %s", oc.id, oc.text)
}

type McQuestion struct {
	abstractQuestion
	opts1 []*Option
}

func NewMcQuestion(text string) *McQuestion {
	mc := &McQuestion{}
	q := NewQuestion(text)
	mc.abstractQuestion = *q
	return mc
}

func (mc *McQuestion) String() string {
	return fmt.Sprintf("McQuestion %s: %s", mc.id, mc.text)
}

type TextQuestion struct {
	abstractQuestion
	//answer string // TODO answer is a -->>fact
}

func NewTextQuestion(text string) *TextQuestion {
	tq := &TextQuestion{}
	q := NewQuestion(text)
	tq.abstractQuestion = *q
	return tq
}

func (tq *TextQuestion) String() string {
	return fmt.Sprintf("TextQuestion %s: %s", tq.id, tq.text)
}

type NumQuestion struct {
	abstractQuestion
	//answer int // TODO answer is a -->>fact
}

func NewNumQuestion(text string) *NumQuestion {
	nq := &NumQuestion{}
	q := NewQuestion(text)
	nq.abstractQuestion = *q
	return nq
}

func (nq *NumQuestion) String() string {
	return fmt.Sprintf("NumQuestion %s: %s", nq.id, nq.text)
}

type ThesisGroup struct {
	AbstractKnowledgebaseObject
	Theses map[string]Thesis
	text   string
}

func NewThesisGroup(text string) *ThesisGroup {
	tg := &ThesisGroup{text: text}
	tg.id = GenerateID(tg)
	return tg
}

func (tg *ThesisGroup) AddThesis(t Thesis) Thesis {
	tt, exists := tg.Theses[t.id]
	if exists == false {
		tg.Theses[t.id] = t
		tt = t
	}
	return tt
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

package model

import (
	"fmt"
	"xpt/tools"
	// "math/big"
)

type Knowledgebase struct {
	text                string
	QuestionGroups      map[string]*QuestionGroup
	allQuestions        map[string]Question
	allTheses           map[string]Thesis
	ThesisGroups        map[string]*ThesisGroup
	StartupQuestions    *QuestionGroup
	AllKnowledgeObjects map[string]*KnowledgebaseObject
}

func NewKnowledgebase(text string) *Knowledgebase {
	k := &Knowledgebase{text: text}
	k.QuestionGroups = make(map[string]*QuestionGroup)
	k.allQuestions = make(map[string]Question)
	k.allTheses = make(map[string]Thesis)
	k.ThesisGroups = make(map[string]*ThesisGroup)
	k.AllKnowledgeObjects = make(map[string]*KnowledgebaseObject)

	return k
}

func (kb *Knowledgebase) describe(sep string) string {
	result := fmt.Sprintf("Knowledgebase: %s", kb.text)
	for _, v := range kb.QuestionGroups {
		result = result + "\n" + sep + v.describe(sep+sep)
	}
	return result
}

func (kb *Knowledgebase) String() string {
	return fmt.Sprintf("Knowledgebase: %s", kb.text)
}

func (kb Knowledgebase) AddQuestionGroup(groups ...*QuestionGroup) *QuestionGroup {
	var first *QuestionGroup
	for _, qg := range groups {
		qt, exists := kb.QuestionGroups[qg.id]
		if exists == false {
			kb.QuestionGroups[qg.id] = qg
			qt = qg
		}
		if first == nil {
			first = qt
		}
	}
	return first
}

func (kb Knowledgebase) AddStartupQuestions(qg *QuestionGroup) *QuestionGroup {
	sq := kb.AddQuestionGroup(qg)
	kb.StartupQuestions = sq

	return sq
}

func (kb Knowledgebase) AddThesisGroup(groups ...*ThesisGroup) *ThesisGroup {
	var first *ThesisGroup
	for _, tg := range groups {
		tt, exists := kb.ThesisGroups[tg.id]
		if exists == false {
			kb.ThesisGroups[tg.id] = tg
			tt = tg
		}
		if first == nil {
			first = tt
		}
	}
	return first
}

func (kb Knowledgebase) AddThesis(tg *ThesisGroup, t *Thesis) {
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
	case *ThesisGroup:
		key = "Q"
	case *QuestionGroup:
		key = "QG"
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

type QuestionGroup struct {
	AbstractKnowledgebaseObject
	Questions map[string]*Question
	text      string
}

func NewQuestionGroup(text string) *QuestionGroup {
	qg := &QuestionGroup{text: text, Questions: make(map[string]*Question)}
	qg.id = GenerateID(qg)
	return qg
}

func (question *QuestionGroup) String() string {
	return fmt.Sprintf("Questiongroup %s: %s", question.id, question.text)
}

func (qg *QuestionGroup) describe(sep string) string {
	result := fmt.Sprintf("%s", qg)
	for _, v := range qg.Questions {
		result = result + "\n" + sep + v.describe(sep+sep)
	}
	return result
}

func (qg *QuestionGroup) AddQuestion(questions ...*Question) *Question {
	var first *Question
	for _, q := range questions {

		qt, exists := qg.Questions[q.Id()]
		if exists == false {
			qg.Questions[q.Id()] = q
			qt = q
		}
		if first == nil {
			first = qt
		}
	}
	return first
}

type Evidence interface {
	String() string
	Id() string
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

func (qg *Question) describe(sep string) string {
	result := fmt.Sprintf("%s", qg)
	/*for k, v := range qg.Question {
		result = result + v.describe()
	}*/
	return result
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

type ThesisGroup struct {
	AbstractKnowledgebaseObject
	Theses map[string]*Thesis
	text   string
}

func NewThesisGroup(text string) *ThesisGroup {
	tg := &ThesisGroup{text: text, Theses: make(map[string]*Thesis)}
	tg.id = GenerateID(tg)
	return tg
}

func (tg *ThesisGroup) AddThesis(theses ...*Thesis) *Thesis {
	var first *Thesis
	for _, t := range theses {
		tt, exists := tg.Theses[t.id]
		if exists == false {
			tg.Theses[t.id] = t
			tt = t
		}
		if first == nil {
			first = tt
		}
	}
	return first
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

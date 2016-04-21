package model

import (
	"fmt"
	"xpt/tools"
)

type Knowledgebase struct {
	text                string
	questionGroups      map[string]*QuestionGroup
	allQuestions        map[string]*Question
	allTheses           map[string]*Thesis
	thesisGroups        map[string]*ThesisGroup
	startupQuestions    *QuestionGroup
	allKnowledgeObjects map[string]KnowledgebaseObject
}

func NewKnowledgebase(text string) *Knowledgebase {
	k := &Knowledgebase{text: text}
	k.questionGroups = make(map[string]*QuestionGroup)
	k.allQuestions = make(map[string]*Question)
	k.allTheses = make(map[string]*Thesis)
	k.thesisGroups = make(map[string]*ThesisGroup)
	k.allKnowledgeObjects = make(map[string]KnowledgebaseObject)

	return k
}

func (kb *Knowledgebase) Describe(sep string) string {
	result := fmt.Sprintf("Knowledgebase: %s", kb.text)
	for _, v := range kb.questionGroups {
		result = result + "\n" + sep + v.Describe(sep+sep)
	}
	for _, v := range kb.thesisGroups {
		result = result + "\n" + sep + v.Describe(sep+sep)
	}
	return result
}

func (kb *Knowledgebase) String() string {
	return fmt.Sprintf("Knowledgebase: %s", kb.text)
}

func (kb *Knowledgebase) AddKnowledbaseObject(objects ...KnowledgebaseObject) KnowledgebaseObject {
	var first KnowledgebaseObject
	for _, o := range objects {
		ot, exists := kb.allKnowledgeObjects[o.Id()]
		if exists == false {
			kb.allKnowledgeObjects[o.Id()] = o
			ot = o

		}
		if first == nil {
			first = ot
		}
	}
	return first
}

func (kb Knowledgebase) AddQuestionGroup(groups ...*QuestionGroup) *QuestionGroup {
	var first *QuestionGroup
	for _, qg := range groups {
		qt, exists := kb.questionGroups[qg.id]
		if exists == false {
			kb.AddKnowledbaseObject(qg)
			kb.questionGroups[qg.id] = qg
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
	kb.startupQuestions = sq

	return sq
}

func (kb Knowledgebase) AddThesisGroup(groups ...*ThesisGroup) *ThesisGroup {
	var first *ThesisGroup
	for _, tg := range groups {
		tt, exists := kb.thesisGroups[tg.id]
		if exists == false {
			kb.AddKnowledbaseObject(tg)
			kb.thesisGroups[tg.id] = tg
			tt = tg
		}
		if first == nil {
			first = tt
		}
	}
	return first
}

type KnowledgebaseObject interface {
	Id() string
}

type AbstractKnowledgebaseObject struct {
	id string
}

func (ako *AbstractKnowledgebaseObject) String() string {
	return fmt.Sprintf("AbstractKnowledgebaseObject", ako.id)
}

func (ako AbstractKnowledgebaseObject) Id() string {
	return ako.id
}

func GenerateID(i interface{}) string {
	var key string
	switch i.(type) {
	default:
		key = "K"
	case *ThesisGroup:
		key = "TG"
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
	Questions map[string]Evidence
	text      string
}

func NewQuestionGroup(text string) *QuestionGroup {
	qg := &QuestionGroup{text: text, Questions: make(map[string]Evidence)}
	qg.id = GenerateID(qg)
	return qg
}

func (question *QuestionGroup) String() string {
	return fmt.Sprintf("Questiongroup %s: %s", question.id, question.text)
}

func (qg *QuestionGroup) Describe(sep string) string {
	result := fmt.Sprintf("%s", qg)
	for _, v := range qg.Questions {
		result = result + "\n" + sep + v.Describe(sep+sep)
	}
	return result
}

func (kb *Knowledgebase) AddQuestion(qg *QuestionGroup, questions ...Evidence) Evidence {
	var first Evidence
	for _, q := range questions {
		qt, exists := qg.Questions[q.Id()]
		if exists == false {
			qg.Questions[q.Id()] = q
			qt = q
			kb.AddKnowledbaseObject(q)
		}
		if first == nil {
			first = qt
		}
	}
	return first
}

func (kb *Knowledgebase) FindObject(id string) KnowledgebaseObject {
	return kb.allKnowledgeObjects[id]
}

type Evidence interface {
	String() string
	Describe(string) string
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

func (qg *Question) Describe(sep string) string {
	result := fmt.Sprintf("%s", qg)
	/*for k, v := range qg.Question {
		result = result + v.Describe()
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

func (theses *ThesisGroup) String() string {
	return fmt.Sprintf("Thesesgroup %s: %s", theses.id, theses.text)
}

func NewThesisGroup(text string) *ThesisGroup {
	tg := &ThesisGroup{text: text, Theses: make(map[string]*Thesis)}
	tg.id = GenerateID(tg)
	return tg
}

func (tg *ThesisGroup) Describe(sep string) string {
	result := fmt.Sprintf("%s", tg)
	for _, v := range tg.Theses {
		result = result + "\n" + sep + v.Describe(sep+sep)
	}
	return result
}

func (kb *Knowledgebase) AddThesis(tg *ThesisGroup, theses ...*Thesis) *Thesis {
	var first *Thesis
	for _, t := range theses {
		tt, exists := tg.Theses[t.id]
		if exists == false {
			tg.Theses[t.id] = t
			tt = t
			kb.AddKnowledbaseObject(t)
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

func (qg *Thesis) Describe(sep string) string {
	result := fmt.Sprintf("%s", qg)
	/*for k, v := range qg.Question {
		result = result + v.Describe()
	}*/
	return result
}

func (theses *Thesis) String() string {
	return fmt.Sprintf("Theses %s: %s", theses.id, theses.text)
}

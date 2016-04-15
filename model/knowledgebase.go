package model

import (
	"fmt"
	"xpt/tools"
)

type Question struct {
	id   string
	text string
}

func NewQuestion(text string) *Question {
	id := tools.GenSID("Q")
	return &Question{text: text, id: id}
}

func (question *Question) String() string {
	return fmt.Sprintf("Question %s: %s", question.id, question.text)
}

type Theses struct {
	id   string
	text string
}

func NewTheses(text string) *Theses {
	id := tools.GenSID("T")
	return &Theses{text: text, id: id}
}

func (theses *Theses) String() string {
	return fmt.Sprintf("Theses %s: %s", theses.id, theses.text)
}

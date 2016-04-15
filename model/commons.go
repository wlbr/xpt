package model

import (
	"fmt"
	"xpt/tools"
)

type Fact struct {
	id       string
	question *Question
}

func NewFact(question *Question) *Fact {
	id := tools.GenSID("F")
	return &Fact{id: id, question: question}
}

func NewFact(question *Question, id string) *Fact {
	id := tools.GenSID("F")
	return &Fact{id: id, question: question}
}

func (fact *Fact) String() string {
	return fmt.Sprintf("fact %s to question %s", fact.id, fact.question.id)
}

type Analysis struct {
	id     string
	theses Theses
	points int
}

func NewAnalysis(theses Theses) *Analysis {
	id := tools.GenSID("A")
	return &Analysis{theses: theses, id: id, points: 0}
}

func (analysis Analysis) String() string {
	return fmt.Sprintf("Analysis %s to theses %s: Points: %s", analysis.id, analysis.theses.id, analysis.points)
}

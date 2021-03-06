package model

import "fmt"

type Case struct {
	kb *Knowledgebase
}

type Fact struct {
	AbstractKnowledgebaseObject
	id       string
	question *Question
}

func NewFact(question *Question) *Fact {
	f := &Fact{question: question}
	f.id = GenerateID(f)
	return f
}

func (fact *Fact) String() string {
	return fmt.Sprintf("fact %s to question %s", fact.id, fact.question.id)
}

type Analysis struct {
	AbstractKnowledgebaseObject
	thesis Thesis
	points int
}

func NewAnalysis(thesis Thesis) *Analysis {
	a := &Analysis{thesis: thesis, points: 0}
	a.id = GenerateID(a)

	return a
}

func (analysis Analysis) String() string {
	return fmt.Sprintf("Analysis %s to theses %s: Points: %d", analysis.id, analysis.thesis.id, analysis.points)
}

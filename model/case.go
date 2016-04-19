package model

import "fmt"

type Case struct {
}

type Fact struct {
	AbstractKnowledgebaseObject
	id       string
	question *abstractQuestion
}

func NewFact(question *abstractQuestion) *Fact {
	f := &Fact{question: question}
	f.id = GenerateID(f)
	return f
}

func (fact *Fact) String() string {
	return fmt.Sprintf("fact %s to question %s", fact.id, fact.Id)
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
	return fmt.Sprintf("Analysis %s to theses %s: Points: %s", analysis.id, analysis.thesis.id, analysis.points)
}

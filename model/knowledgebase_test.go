package model

import "testing"

var kb = NewKnowledgebase()

func prepareTests() {

}

func TestKB(t *testing.T) {

}

func TestQuestions(t *testing.T) {

	ql := []KnowledgebaseObject{NewQuestion("My first Question"),
		NewTextQuestion("Another Question"),
		NewNumQuestion("Another Question"),
		NewOcQuestion("OneChoice Question"),
		NewMcQuestion("MultipleChoice Question"),
		NewTheses("This is an assumption."),
	}

	assertObject(t, ql[0], ql[0])

}

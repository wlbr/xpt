package model

import (
	"fmt"
	"testing"
)

var kb *Knowledgebase

func prepareTests() {
	kb = NewKnowledgebase("IT Support")

	tpr := NewThesisGroup("Printer")
	tn := NewThesisGroup("Network")

	kb.AddThesisGroup(tn, tpr)

	sqgr := NewQuestionGroup("Startup Questions")
	kb.AddStartupQuestions(sqgr)
	qgrn := NewQuestionGroup("Network setup")
	qgrp := NewQuestionGroup("Printer setup")

	kb.AddQuestionGroup(sqgr, qgrn, qgrp)

	q1 := NewQuestion("Please enter your name.")
	q2 := NewQuestion("Please enter your email address.")
	q3 := NewQuestion("What operating System are you using?")
	q4 := NewQuestion("Do you have a problem with your a) network b) printer?")

	sqgr.AddQuestion(q1, q2, q3, q4)

	q5 := NewQuestion("Printer question 1")
	q6 := NewQuestion("Printer question 2")
	q7 := NewQuestion("Printer question 3")
	qgrp.AddQuestion(q5, q6, q7)

	q8 := NewQuestion("Network question 1")
	q9 := NewQuestion("Network question 2")
	q10 := NewQuestion("Network question 3")
	qgrn.AddQuestion(q8, q9, q10)
}

func TestKB(t *testing.T) {
	prepareTests()

	fmt.Printf(kb.describe(" "))

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

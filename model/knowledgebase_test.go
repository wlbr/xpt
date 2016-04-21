package model

import (
	"fmt"
	"strings"
	"testing"
	"xpt/tools"
)

func prepareTests() *Knowledgebase {
	kb := NewKnowledgebase("IT Support")
	tools.ResetAllSID()

	tpr := NewThesisGroup("Printer")
	tn := NewThesisGroup("Network")

	kb.AddThesisGroup(tn, tpr)
	tpdosx := NewTheses("Please install the PS PPD File for the printer type Canon iR---ADV C5030/5035 at http://www.canon.de/support/products/imagerunner/imagerunner-advance-c5030.aspx?type=drivers&language=&os=OS%20X%2010.11%20(El%20Capitan)")
	tpdwin := NewTheses("Please install the PS PPD File for the printer type Canon iR---ADV C5030/5035 at http://www.canon.de/support/products/imagerunner/imagerunner-advance-c5030.aspx?type=drivers&language=&os=Windows%2010%20(64-bit)")
	kb.AddThesis(tpr, tpdosx, tpdwin)

	sqgr := NewQuestionGroup("Startup Questions")
	kb.AddStartupQuestions(sqgr)
	qgrn := NewQuestionGroup("Network setup")
	qgrp := NewQuestionGroup("Printer setup")

	kb.AddQuestionGroup(sqgr, qgrn, qgrp)

	q1 := NewQuestion("Please enter your name.")
	q2 := NewQuestion("Please enter your email address.")
	q3 := NewQuestion("What operating system are you using?")
	q4 := NewQuestion("Where are you located?")
	q5 := NewOcQuestion("Do you have a problem with your a) network b) printer?")

	kb.AddQuestion(sqgr, q1, q2, q3, q4, q5)

	q6 := NewQuestion("Printer question 1")
	q7 := NewQuestion("Printer question 2")
	q8 := NewQuestion("Printer question 3")
	kb.AddQuestion(qgrp, q6, q7, q8)

	q9 := NewQuestion("Network question 1")
	q10 := NewQuestion("Network question 2")
	q11 := NewQuestion("Network question 3")
	kb.AddQuestion(qgrn, q9, q10, q11)

	return kb
}

func TestKB(t *testing.T) {
	kb := prepareTests()
	fmt.Printf("\n\n\n%s\n\n\n", kb.Describe("  "))
}

func TestAllKnowlegeObjectsList(t *testing.T) {
	kb := prepareTests()
	assertInt(t, 18, len(kb.allKnowledgeObjects))
	fmt.Printf("All objects in %s: %d\n", kb, len(kb.allKnowledgeObjects))
	for _, o := range kb.allKnowledgeObjects {
		fmt.Println(o)
	}
	desc := kb.Describe("  ")

	fmt.Printf("lines %d", strings.Count(desc, "\n"))
	//assertInt(t, strings.Count(desc, "\n"), strings.Count(desc, "\n")+1)
}

func TestFindObject(t *testing.T) {
	kb := prepareTests()
	searchID := "Q3"
	assertObject(t, searchID, kb.FindObject(searchID).Id())
	assertObject(t, "What operating system are you using?", kb.FindObject(searchID).(*Question).text)
	fmt.Println(kb.FindObject(searchID).(*Question).text)
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

func TestDescribeVsString(t *testing.T) {
	kb := prepareTests()
	searchID := kb.FindObject("Q5")

	fmt.Println("ToString: ", searchID)
	fmt.Println("Describe: ", searchID.(*OcQuestion).Describe(""))
}
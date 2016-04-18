package main

import (
	"fmt"
	"xpt/model"
)

func testKnowledgebase() {
	kb := []model.KnowledgebaseObject{model.NewQuestion("My first Question"),
		model.NewTextQuestion("Another Question"),
		model.NewNumQuestion("Another Question"),
		model.NewOcQuestion("OneChoice Question"),
		model.NewMcQuestion("MultipleChoice Question"),
		model.NewTheses("This is an assumption."),
	}

	for _, i := range kb {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println()
	testKnowledgebase()
}

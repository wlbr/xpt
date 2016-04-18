package model

import "fmt"

type Rule struct {
	AbstractKnowledgebaseObject
	question Question
	thesis   Thesis
	//condition Condition
	action Action
}

func (rule *Rule) String() string {
	return fmt.Sprintf("Rule %s: %s", rule.id)
}

/*type Condition struct {
}
*/

type Action struct {
}

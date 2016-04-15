package main

import (
	"fmt"
	"xpt/tools"
)

func testIdGeneration() {

	fmt.Printf("first %s \n", tools.GenSID("first"))
	fmt.Printf("first %s \n", tools.GenSID("first"))
	fmt.Printf("second %s \n", tools.GenSID("second"))
	fmt.Printf("second %s \n", tools.GenSID("second"))
	fmt.Printf("first %s \n", tools.GenSID("first"))
	tools.SetLastSID("first", 10)
	fmt.Print("setting first to 10")
	fmt.Printf("first %s \n", tools.GenSID("first"))
	fmt.Printf("second %s \n", tools.GenSID("second"))
	fmt.Printf("first %s \n", tools.GenSID("first"))
}

func main() {
	testIdGeneration()
}

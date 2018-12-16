package gotils

import (
	"fmt"
	"strconv"
	"testing"
)

func assert(t *testing.T, gold interface{}, test interface{}) {
	// fmt.Printf("gold=%s test=%s\n", gold, test)
	if gold != test {
		fmt.Printf("Assert failed, %v != %v\t", gold, test)
		t.Fail()
	}
}

func cleanup() {
	ResetAllSID()
}

func TestDistinctNamespaces(t *testing.T) {
	first := GenSID("first")
	second := GenSID("first")
	third := GenSID("first")

	fourth := GenSID("second")
	fifth := GenSID("second")

	sixth := GenSID("third")
	seventh := GenSID("first")

	assert(t, "first1", first)
	assert(t, "first2", second)
	assert(t, "first3", third)
	assert(t, "second1", fourth)
	assert(t, "second2", fifth)
	assert(t, "third1", sixth)
	assert(t, "first4", seventh)

	cleanup()
}

func TestSettingTheId(t *testing.T) {
	first := GenSID("fourth")
	second := GenSID("fourth")
	third := GenSID("fourth")

	fourth := GenSID("fifth")
	fifth := GenSID("fifth")

	SetLastSID("fourth", 10)

	sixth := GenSID("sixth")
	seventh := GenSID("fourth")

	assert(t, "fourth1", first)
	assert(t, "fourth2", second)
	assert(t, "fourth3", third)
	assert(t, "fifth1", fourth)
	assert(t, "fifth2", fifth)
	assert(t, "sixth1", sixth)
	assert(t, "fourth11", seventh)

	cleanup()
}

func TestResetSID(t *testing.T) {
	first := GenSID("first")
	GenSID("first")
	GenSID("first")
	second := GenSID("first")
	GenSID("first")
	GenSID("first")
	GenSID("first")
	third := GenSID("first")

	n1, _ := strconv.Atoi(first[5:])
	n2, _ := strconv.Atoi(second[5:])
	n3, _ := strconv.Atoi(third[5:])

	assert(t, 3, n2-n1)
	assert(t, 7, n3-n1)
	assert(t, 4, n3-n2)

	ResetSID("first")
	l1, _ := strconv.Atoi(LastSID("first")[5:])
	assert(t, 0, l1)

	first = GenSID("first")
	GenSID("first")
	second = GenSID("first")
	n1, _ = strconv.Atoi(first[5:])
	n2, _ = strconv.Atoi(second[5:])

	assert(t, 2, n2-n1)

	cleanup()
}

func TestResetAllSID(t *testing.T) {

	first := GenSID("first")
	GenSID("first")
	GenSID("first")
	second := GenSID("first")
	GenSID("first")
	GenSID("first")
	GenSID("first")
	third := GenSID("first")
	fourth := GenSID("second")
	GenSID("second")
	fifth := GenSID("second")
	GenSID("second")
	GenSID("second")
	sixth := GenSID("second")

	n1, _ := strconv.Atoi(first[5:])
	n2, _ := strconv.Atoi(second[5:])
	n3, _ := strconv.Atoi(third[5:])
	n4, _ := strconv.Atoi(fourth[6:])
	n5, _ := strconv.Atoi(fifth[6:])
	n6, _ := strconv.Atoi(sixth[6:])

	assert(t, 3, n2-n1)
	assert(t, 7, n3-n1)
	assert(t, 4, n3-n2)

	assert(t, 2, n5-n4)
	assert(t, 5, n6-n4)
	assert(t, 3, n6-n5)

	ResetAllSID()

	first = GenSID("first")
	GenSID("first")
	second = GenSID("first")
	n1, _ = strconv.Atoi(first[5:])
	n2, _ = strconv.Atoi(second[5:])
	assert(t, 2, n2-n1)

	l1, _ := strconv.Atoi(LastSID("first")[5:])
	l2, _ := strconv.Atoi(LastSID("second")[6:])
	assert(t, 3, l1)
	assert(t, 0, l2)

	cleanup()
}

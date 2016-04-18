package tools

import "testing"

func assert(t *testing.T, gold string, test string) {
	// fmt.Printf("gold=%s test=%s\n", gold, test)
	if gold != test {
		t.Fail()
	}
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
}

package tinycache

import (
	"testing"
)

func BenchmarkAddImport(b *testing.B) {
	testCache := NewCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			testCache.Add(getRandomElement())
		}
	}
}

func TestMaxElements(t *testing.T) {
	testCache := NewCache()
	testCache.SetMaxElements(10)

	for i := 0; i < 5; i++ {
		t.Logf("adding element %d", i)
		testCache.Add(getRandomElement())
	}

	if testCache.Count() != 5 {
		t.Errorf("expected number of elements: %d, have %d", 5, testCache.Count())
	}
}

func TestCapTo5(t *testing.T) {
	testCache := NewCache()
	testCache.SetMaxElements(5)

	for i := 0; i < 10; i++ {
		t.Logf("adding element %d", i)
		testCache.Add(getRandomElement())
	}

	if testCache.Count() != 5 {
		t.Errorf("expected number of elements: %d, have %d", 5, testCache.Count())
	}
}

func getRandomElement() interface{} {
	element := struct {
		id, age             int
		firstName, lastName string
	}{
		age:       37,
		firstName: "foo",
		lastName:  "bar",
	}

	return element
}

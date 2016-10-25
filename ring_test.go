package ring

import (
	"testing"
)

func TestNew(t *testing.T) {
	n := New(5)
	if n.Size != 5 {
		t.Errorf("New failed")
	}
}

func TestGetEqualSet(t *testing.T) {
	r := New(5)
	for i := 0; i < 10; i++ {
		r.Set(i)
		if r.Get() != i {
			t.Errorf("Get or Set buffer error")
		}
	}
}

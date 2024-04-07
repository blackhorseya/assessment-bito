package rbtree

import (
	"testing"
)

func TestInsertAndDelete(t *testing.T) {
	rbt := New()

	m := 0
	n := 1000
	for m < n {
		rbt.Insert(Int(m))
		m++
	}
	if rbt.Len() != uint(n) {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), n)
	}

	for m > 0 {
		rbt.Delete(Int(m))
		m--
	}
	if rbt.Len() != 1 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 1)
	}
}

type testVar struct {
	Height uint
	ID     string
}

func (t *testVar) Less(than Item) bool {
	return t.Height < than.(*testVar).Height
}

func TestGet(t *testing.T) {
	rbt := New()

	rbt.Insert(&testVar{
		Height: 170,
		ID:     "1",
	})

	got := rbt.Get(&testVar{
		Height: 170,
		ID:     "1",
	})

	if got == nil {
		t.Errorf("tree.Get() = nil, expect not nil")
	}

	rbt.Delete(&testVar{
		Height: 170,
		ID:     "1",
	})

	got = rbt.Get(&testVar{
		Height: 170,
		ID:     "1",
	})

	if got != nil {
		t.Errorf("tree.Get() = %v, expect nil", got)
	}
}

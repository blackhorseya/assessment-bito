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

package rbtree

import (
	"reflect"
	"testing"
)

func TestDescend(t *testing.T) {
	rbt := New()

	m := 0
	n := 10
	for m < n {
		rbt.Insert(Int(m))
		m++
	}

	var ret []Item

	rbt.Descend(Int(1), func(i Item) bool {
		ret = append(ret, i)
		return true
	})
	expected := []Item{Int(1), Int(0)}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}

	ret = nil
	rbt.Descend(Int(10), func(i Item) bool {
		ret = append(ret, i)
		return true
	})
	expected = []Item{Int(9), Int(8), Int(7), Int(6), Int(5), Int(4), Int(3), Int(2), Int(1), Int(0)}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}

func TestAscend(t *testing.T) {
	rbt := New()

	rbt.Insert(String("a"))
	rbt.Insert(String("b"))
	rbt.Insert(String("c"))
	rbt.Insert(String("d"))

	rbt.Delete(rbt.Min())

	var ret []Item
	rbt.Ascend(rbt.Min(), func(i Item) bool {
		ret = append(ret, i)
		return true
	})

	expected := []Item{String("b"), String("c"), String("d")}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}

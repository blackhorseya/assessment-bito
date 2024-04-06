package rbtree

// Iterator is the function of iteration entity which would be
// used by those functions like `Ascend`, `Dscend`, etc.
//
// A typical Iterator with Print :
//
//	func loop_with_print(item rbtree.Item) bool {
//	        i, ok := item.(XXX)
//	        if !ok {
//	                return false
//	        }
//	        fmt.Printf("%+v\n", i)
//	        return true
//	}
type Iterator func(i Item) bool

// Ascend will call iterator once for each element greater or equal than pivot
// in ascending order. It will stop whenever the iterator returns false.
func (t *RBTree) Ascend(pivot Item, iterator Iterator) {
	t.ascend(t.root, pivot, iterator)
}

func (t *RBTree) ascend(x *Node, pivot Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(x.Item, pivot) {
		if !t.ascend(x.Left, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.ascend(x.Right, pivot, iterator)
}

// Descend will call iterator once for each element less or equal than pivot
// in descending order. It will stop whenever the iterator returns false.
func (t *RBTree) Descend(pivot Item, iterator Iterator) {
	t.descend(t.root, pivot, iterator)
}

func (t *RBTree) descend(x *Node, pivot Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(pivot, x.Item) {
		if !t.descend(x.Right, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.descend(x.Left, pivot, iterator)
}

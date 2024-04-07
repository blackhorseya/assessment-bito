package rbtree

// Len returns the number of items in the tree.
func (t *RBTree) Len() uint {
	return t.count
}

// Get returns the item in the tree that is equal to the given item.
func (t *RBTree) Get(item Item) Item {
	if item == nil {
		return nil
	}

	// The `color` field here is nobody
	ret := t.search(&Node{t.NIL, t.NIL, t.NIL, RED, item})
	if ret == nil {
		return nil
	}

	return ret.Item
}

// Insert adds an item to the tree.
func (t *RBTree) Insert(item Item) {
	if item == nil {
		return
	}

	t.insert(&Node{
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: t.NIL,
		Color:  RED,
		Item:   item,
	})
}

// Delete removes an item from the tree.
func (t *RBTree) Delete(item Item) Item {
	if item == nil {
		return nil
	}

	return t.delete(&Node{
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: t.NIL,
		Color:  RED,
		Item:   item,
	}).Item
}

// Min returns the smallest item in the tree.
func (t *RBTree) Min() Item {
	x := t.min(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}

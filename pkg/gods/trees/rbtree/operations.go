package rbtree

// Len returns the number of items in the tree.
func (t *RBTree) Len() uint {
	return t.count
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

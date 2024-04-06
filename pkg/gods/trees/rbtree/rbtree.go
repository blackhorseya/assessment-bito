package rbtree

// Item is the interface that the tree will use to store items.
type Item interface {
	Less(than Item) bool
}

const (
	ColorRed   = false
	ColorBlack = true
)

// Node is a single element within the tree.
type Node struct {
	Item

	Left   *Node
	Right  *Node
	Parent *Node
	Color  bool
}

// RBTree is a red-black tree.
type RBTree struct {
	NIL *Node

	root  *Node
	count uint
}

// New returns an initialized red-black tree.
func New() *RBTree {
	node := &Node{
		Item:   nil,
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Color:  true,
	}

	return &RBTree{
		NIL:   node,
		root:  node,
		count: 0,
	}
}

func less(a, b Item) bool {
	return a.Less(b)
}

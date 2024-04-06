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

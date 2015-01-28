// Left-leaning Red-black Tree
// http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
// http://www.cs.princeton.edu/~rs/talks/LLRB/RedBlack.pdf
// Based on 2-3 Trees.
//
// 3-nodes are represented as a node with a left Red child.
// 3-nodes are left-leaning.
//
// 4-nodes are represented as a node with two Red children.
//
// Disallowed (invariants):
//  1. Right-leaning 3-node representation
//  2. Two reds in a row
//
// Current operations:
//  Max, Min, DeleteMin, DeleteMax, Get, Insert, Delete
//
package llrb

// Color specifies the color of a node.
type Color bool

const (
	Red   Color = false
	Black Color = true
)

// Key of a Node.
type Key interface {
	// Less returns true is a < the receiver element.
	Less(a interface{}) bool
}

// Value of a Node.
type Value interface{}

// Node represents a node in the LLRB Tree.
type Node struct {
	Key
	Value
	Color

	left, right *Node
}

// Tree represents a 2-3 LLRB Tree.
type Tree struct {
	root *Node
}

// New creates a new LLRB Tree.
func New() *Tree {
	return &Tree{
		root: nil,
	}
}

// Get searches for a Node is the Tree based on a key. If the node is
// found Node.Value is returned, otherwise nil.
func (t *Tree) Get(k Key) Value {
	return get(t.root, k)
}

func get(n *Node, k Key) interface{} {
	for n != nil {
		if k == n.Key {
			return n.Value
		} else if k.Less(n.Key) {
			n = n.left
		} else {
			n = n.right
		}
	}
	return nil
}

// Len returns the number of Nodes in the Tree.
func (t *Tree) Len() int {
	return len(t.root)

}

func len(n *Node) int {
	if n == nil {
		return 0
	}
	return len(n.left) + len(n.right) + 1
}

// Height returns the height of the Tree. The height is the number
// of levels before the bottom most node is found.
func (t *Tree) Height() int {
	return height(t.root) + 1
}

func height(n *Node) int {
	if n == nil {
		return 0
	}

	leftHeight := height(n.left)
	rightHeight := height(n.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Insert inserts a new node with Key = k and Value = v.
// If k already exists, the current value with be replace with v.
func (t *Tree) Insert(k Key, v Value) {
	t.root = insert(t.root, k, v)
	if t.root != nil {
		t.root.Color = Black
	}
}

// insert inserts a node into the Tree as it would in a regular BST.
// After the insertion has been completed if an invariant has been violated
// it will be fixed, maintaining O(log n) height.
func insert(n *Node, k Key, v Value) *Node {
	if n == nil {
		return &Node{
			Key:   k,
			Value: v,
			Color: Red,
		}
	}

	// If colors are flipped here this turns into a 2-3-4 Tree.

	if k == n.Key {
		n.Value = v
	} else if k.Less(n.Key) {
		n.left = insert(n.left, k, v)
	} else {
		n.right = insert(n.right, k, v)
	}

	n = fixUp(n)

	return n
}

// Delete deletes the node where the Key == k.
func (t *Tree) Delete(k Key) {
	t.root = delete(t.root, k)
	if t.root != nil {
		t.root.Color = Black
	}
}

func delete(n *Node, k Key) *Node {
	if n == nil {
		return nil
	}

	if k.Less(n.Key) {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = delete(n.left, k)
	} else {
		if isRed(n.left) {
			n = rotateRight(n)
		}

		if k == n.Key && n.right == nil {
			return nil
		}

		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}

		// Found node switch K/V with min node of right child.
		// Delete min node of the right child.
		if k == n.Key {
			mk, mv := min(n.right)
			n.Key = mk
			n.Value = mv
			n.right = deleteMin(n.right)
		} else {
			n.right = delete(n.right, k)
		}

	}

	return fixUp(n)
}

// for delete routine
func min(n *Node) (Key, Value) {
	if n == nil {
		return nil, nil
	}
	for n.left != nil {
		n = n.left
	}

	return n.Key, n.Value
}

// DeleteMin deletes the minimum element of the Tree.
func (t *Tree) DeleteMin() {
	t.root = deleteMin(t.root)
	if t.root != nil {
		t.root.Color = Black
	}
}

// Min returns the minimum Key.
func (t *Tree) Min() Key {
	n := t.root
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}

	return n.Key
}

// Max returns the maximum Key.
func (t *Tree) Max() Key {
	n := t.root
	if n == nil {
		return nil
	}

	for n.right != nil {
		n = n.right
	}

	return n.Key
}

func deleteMin(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return nil
	}

	if !isRed(n.left) && !isRed(n.left.left) {
		n = moveRedLeft(n)
	}

	n.left = deleteMin(n.left)

	return fixUp(n)
}

// DeleteMax deletes the minimum element of the Tree
func (t *Tree) DeleteMax() {
	t.root = deleteMax(t.root)
	if t.root != nil {
		t.root.Color = Black
	}
}

func deleteMax(n *Node) *Node {
	if n == nil {
		return nil
	}

	if isRed(n.left) {
		n = rotateRight(n)
	}

	if n.right == nil {
		return nil
	}

	if !isRed(n.right) && !isRed(n.right.left) {
		n = moveRedRight(n)
	}

	n.right = deleteMax(n.right)

	return fixUp(n)
}

func rotateRight(n *Node) *Node {
	x := n.left
	n.left = x.right
	x.right = n
	x.Color = n.Color
	n.Color = Red
	return x
}

func rotateLeft(n *Node) *Node {
	x := n.right
	n.right = x.left
	x.left = n
	x.Color = n.Color
	n.Color = Red
	return x
}

func moveRedLeft(n *Node) *Node {
	colorFlip(n)
	if isRed(n.right.left) {
		n.right = rotateRight(n.right)
		n = rotateLeft(n)
		colorFlip(n)
	}
	return n
}

func moveRedRight(n *Node) *Node {
	colorFlip(n)
	if isRed(n.left.left) {
		n = rotateRight(n)
		colorFlip(n)
	}
	return n
}

// isRed returns true if n.Color == Red, false otherwise.
func isRed(n *Node) bool {
	if n == nil {
		return false
	}
	return n.Color == Red
}

// colorFlip flips the Color of n and its direct children.
func colorFlip(n *Node) {
	n.Color = !n.Color
	n.left.Color = !n.left.Color
	n.right.Color = !n.right.Color

}

// fixUp fixes any violated invariants on Node n.
func fixUp(n *Node) *Node {
	// Disallowed right-leaning.
	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}

	// Disallowed two reds in a row.
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}

	// Change if 4-node.
	if isRed(n.left) && isRed(n.right) {
		colorFlip(n)
	}
	return n
}

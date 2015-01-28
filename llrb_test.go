package llrb_test

import (
	"math/rand"
	"testing"

	"github.com/domluna/go-llrb"
)

func TestBasics(t *testing.T) {
	tree := llrb.New()

	v := tree.Get(llrb.IntKey(5))
	if v != nil {
		t.Errorf("Get(llrb.IntKey(5)): got %v, want nil", v)
	}

	v = tree.Max()
	if v != nil {
		t.Errorf("Max() :got %v, want nil", v)
	}
	v = tree.Min()
	if v != nil {
		t.Errorf("Min(): got %v, want nil", v)
	}

	tree.Insert(llrb.IntKey(10), "root")
	tree.Insert(llrb.IntKey(5), "min")
	tree.Insert(llrb.IntKey(15), "max")

	max := tree.Max()
	if max != llrb.IntKey(15) {
		t.Errorf("Max(): got %v, want %v", max, 15)
	}
	min := tree.Min()
	if min != llrb.IntKey(5) {
		t.Errorf("Min(): got %v, want %v", min, 5)
	}

	tree.DeleteMin()
	min = tree.Min()
	if min != llrb.IntKey(10) {
		t.Errorf("DeleteMin(): got %v, want %v", min, 10)
	}

	tree.DeleteMax()
	max = tree.Max()
	if max != llrb.IntKey(10) {
		t.Errorf("DeleteMax(): got %v, want %v", max, 10)
	}

	tree.DeleteMax()
	tree.DeleteMin() // in here for good measure

	t.Logf("Tree is empty we should get nil for Max()/Min()")

	v = tree.Max()
	if v != nil {
		t.Errorf("Min(): got %v, want nil", v)
	}

	v = tree.Min()
	if v != nil {
		t.Errorf("Max(): got %v, want nil", v)
	}
}

func TestInsertAndGet(t *testing.T) {
	tree := llrb.New()
	nNodes := 1 << 16
	keys := rand.Perm(nNodes)

	for _, k := range keys {
		tree.Insert(llrb.IntKey(k), k)
	}

	for _, k := range keys {
		if v := tree.Get(llrb.IntKey(k)); v == nil {
			t.Errorf("Get: %v, got %v, want %v", k, nil, k)
		}
	}

}

func TestHeight(t *testing.T) {
	tree := llrb.New()
	nNodes := 1 << 16

	for i := 0; i < nNodes; i++ {
		n := rand.Int()
		tree.Insert(llrb.IntKey(n), n)
	}

	h := tree.Height()
	if h >= 2*h {
		t.Errorf("Height: should be <= 2 * %d, got %d", 16, h)
	}
}

func TestDelete(t *testing.T) {
	// t.Skip()
	tree := llrb.New()
	nNodes := 1 << 16
	keys := rand.Perm(nNodes)

	for _, k := range keys {
		tree.Insert(llrb.IntKey(k), k)
	}

	l := tree.Len()
	if l != nNodes {
		t.Errorf("Len: got %d, want %d", l, nNodes)
	}

	// Delete
	for _, k := range keys {
		tree.Delete(llrb.IntKey(k))
	}

	l = tree.Len()
	if l != 0 {
		t.Errorf("Len: got %d, want %d", l, 0)
	}

	// Delete again for good measure
	tree.Delete(llrb.IntKey(10))
}

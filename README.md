LLRB
----

2-3 LLRB Tree in Go.

This started off as a deep dive into 2-3 and 2-3-4 trees which then lead to
LLRB trees. The following implementation is largely dervied from these notes.

1. http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
2. http://www.cs.princeton.edu/~rs/talks/LLRB/RedBlack.pdf

Red Black Trees (RBT) are a self balancing binary tree, therefore all operations are O(log n). 2-3 LLRB Trees have the nice benefit that the height is at most 2 * logn. From my experiments it's ~1.5 * logn.

Due to the sorted property, BBST can/are used to implement various data structures. These includes TreeMaps, Queues, etc.

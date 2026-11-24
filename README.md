# Floor in a BST

> Striver SDE Sheet 138 · Day 20 BST — one descent: smaller nodes are candidates before moving right, larger ones just steer left. Go, stdlib only, tested.

## Problem

The floor of a key in a BST: the largest stored value ≤ the key, with a clean "doesn't exist" answer when the key is below the whole tree.

## Approach

A single steering descent with one memory slot. A node equal to the key *is* the floor — return. A node smaller than the key is a valid floor candidate, but a better (larger, still ≤ key) one can only be in its right subtree — record and go right. A node larger than the key can't be the floor and neither can its right side — go left. The recorded candidate only improves along the walk, so the last one is the answer. Returning `(value, ok)` instead of a sentinel keeps "no floor" unambiguous even when values include every int.

**Complexity:** O(h) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).

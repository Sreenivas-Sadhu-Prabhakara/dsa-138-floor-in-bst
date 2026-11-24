package main

import "testing"

func TestFloor(t *testing.T) {
	root := tn(8, tn(4, tn(2, nil, nil), tn(6, nil, nil)), tn(12, tn(10, nil, nil), tn(14, nil, nil)))
	cases := []struct {
		key, want int
		ok        bool
	}{
		{11, 10, true},
		{8, 8, true},  // exact hit
		{1, 0, false}, // below the minimum
		{15, 14, true},
		{5, 4, true},
		{2, 2, true},
	}
	for _, c := range cases {
		got, ok := Floor(root, c.key)
		if ok != c.ok || (ok && got != c.want) {
			t.Errorf("Floor(%d) = (%d,%v), want (%d,%v)", c.key, got, ok, c.want, c.ok)
		}
	}
	if _, ok := Floor(nil, 5); ok {
		t.Error("empty tree has no floor")
	}
}

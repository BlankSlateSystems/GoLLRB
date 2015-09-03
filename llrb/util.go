// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Substantially changed by Gabe (with apologies to Petar) to decouple
// item types from the comparison function.

package llrb

type Int int

type Float32 float32

type String string

// Use in LLRB.New constructor to make a tree that assumes integer items,
// and sorts from smallest to largest.
func NaturalSortLessInt(a, b interface{}) bool {
	return a.(Int) < b.(Int)
}

// Use in LLRB.New constructor to make a tree that assumes float32 items,
// and sorts from smallest to largest.
func NaturalSortLessFloat(a, b interface{}) bool {
	return a.(Float32) < b.(Float32)
}

// Use in LLRB.New constructor to make a tree that assumes string items,
// and sorts alphabetically from smallest to largest (e.g. "a" < "b" "c".
func NaturalSortLessString(a, b interface{}) bool {
	return a.(String) < b.(String)
}

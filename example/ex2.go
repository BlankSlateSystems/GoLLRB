package main

import (
	"fmt"
	"github.com/blankslatesystems/GoLLRB/llrb"
)

type Float float32

func (x Float) Less(than llrb.Item) bool {
	// For an expression 'than' of interface type and a type Float,
	// the expression than.(Float) asserts that 'than' is not nil
	// and that the value stored in 'than' is of type 'Float'.
	//
	// This is called a type assertion.
	//
	// In other words, it essentially casts 'than' to type 'Float',
	// and has the value of than as cast to Float. Clear as mud?
	return x < than.(Float)
}

func Print(item llrb.Item) bool {
	i, ok := item.(Float) // i, ok := item.(llrb.Int)
	if !ok {
		return false
	}
	fmt.Printf("%f of type %T\n", float32(i), i) // fmt.Println(int(i))
	return true
}

func main() {
	tree := llrb.New()
	// replace llrb.Int(n) with Float(n)
	tree.ReplaceOrInsert(Float(1))
	tree.ReplaceOrInsert(Float(2))
	tree.ReplaceOrInsert(Float(3))
	tree.ReplaceOrInsert(Float(4))
	tree.DeleteMin()
	tree.Delete(Float(4))
	tree.AscendGreaterOrEqual(tree.Min(), Print)
}

package main

import (
	"fmt"
	"github.com/blankslatesystems/GoLLRB/llrb"
)

func Print2(item llrb.Item) bool {
	i, ok := item.(llrb.Float32) // i, ok := item.(llrb.Int)
	if !ok {
		return false
	}
	fmt.Printf("%f of type %T\n", float32(i), i) // fmt.Println(int(i))
	return true
}

// func main() {
// 	tree := llrb.New(llrb.NaturalSortLessFloat)
// 	// replace llrb.Int(n) with Float(n)
// 	tree.ReplaceOrInsert(llrb.Float32(1))
// 	tree.ReplaceOrInsert(llrb.Float32(2))
// 	tree.ReplaceOrInsert(llrb.Float32(3))
// 	tree.ReplaceOrInsert(llrb.Float32(4))
// 	tree.DeleteMin()
// 	tree.Delete(llrb.Float32(4))
// 	tree.AscendGreaterOrEqual(tree.Min(), Print2)
// }

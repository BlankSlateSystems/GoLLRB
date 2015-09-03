package main

import (
	"fmt"
	"github.com/blankslatesystems/GoLLRB/llrb"
	"math/rand"
	"time"
)

type Pt struct {
	x, y float32
}

func (p Pt) Less(other llrb.Item) bool {
	return p.x < other.(Pt).x
}

func Shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func Print(item llrb.Item) bool {
	i, ok := item.(Pt)
	if !ok {
		return false
	}
	fmt.Printf("%f,%f of type %T\n", i.x, i.y, i)
	return true
}

func main() {
	tree := llrb.New()
	points := []Pt{}
	numPoints := 100000
	for i := 0; i < numPoints; i++ {
		points = append(points, Pt{rand.Float32(), rand.Float32()})
	}
	startInsert := time.Now()
	for _, dasPt := range points {
		tree.ReplaceOrInsert(dasPt)
	}
	elapsedInsert := time.Since(startInsert)

	pointIndexes := make([]int, numPoints)
	for i := 0; i < numPoints; i++ {
		pointIndexes[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	Shuffle(pointIndexes)
	startDelete := time.Now()
	for _, v := range pointIndexes {
		tree.Delete(points[v])
	}
	elapsedDelete := time.Since(startDelete)
	fmt.Printf("INS %d points in %v\n", len(points), elapsedInsert)
	fmt.Printf("DEL %d points in %v\n", len(points), elapsedDelete)
}

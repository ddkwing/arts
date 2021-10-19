package algorithm

import (
	"ddkwing/arts/algorithm/structures"
	"fmt"
	"testing"
)

type question0141 struct {
	para0141 []int
	ans      bool
}

func Test_Problem0141(T *testing.T) {
	q := question0141{
		para0141: []int{3, 2, 0, -4},
		ans:      false,
	}
	fmt.Printf("[input]: %v,[output]: %v", q, hasCycle(structures.Ints2List(q.para0141)))
}

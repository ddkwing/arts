package algorithm

import (
	"fmt"
	"testing"
)

type question26 struct {
	para26 []int
	ans26  int
}

func Test_Problem26(T *testing.T) {
	q := question26{
		para26: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		ans26:  5,
	}
	fmt.Printf("[input]: %v, [output]: %v", q, removeDuplicates(q.para26))
}

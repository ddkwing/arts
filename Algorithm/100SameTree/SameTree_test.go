package algorithm

import (
	"ddkwing/arts/algorithm/structures"
	"fmt"
	"testing"
)

type question100 struct {
	para para100
	ans  bool
}

type para100 struct {
	s []int
	t []int
}

func Test_Problem100(T *testing.T) {
	q := question100{
		para100{
			[]int{1, 3, 3},
			[]int{1, 2, 3},
		},
		false,
	}

	s := structures.Ints2TreeNode(q.para.s)
	t := structures.Ints2TreeNode(q.para.t)
	fmt.Printf("[input]:%v,[output]:%v", q, isSameTree(s, t))

}

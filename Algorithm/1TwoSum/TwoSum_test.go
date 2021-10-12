package algorithm

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	num    []int
	target int
}

type ans1 struct {
	num []int
}

func Test_Problem1(T *testing.T) {
	q := question1{
		para1{[]int{1, 2, 4, 7}, 11},
		ans1{[]int{1, 2}},
	}
	fmt.Printf("【输入】：%v, 【输出】：%v\n", q, twoSum(q.para1.num, q.para1.target))
}

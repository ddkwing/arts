package algorithm

import (
	"fmt"
	"testing"
)

type question2 struct {
	para2 string
	ans2  int
}

func Test_Problem2(T *testing.T) {
	q := question2{
		para2: "pwwkew",
		ans2:  3,
	}
	fmt.Printf("【问题】%v,【输出】%v\n", q, lengthOfLongestString(q.para2))
}

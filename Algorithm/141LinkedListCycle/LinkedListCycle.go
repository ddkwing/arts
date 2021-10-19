package algorithm

import (
	"ddkwing/arts/algorithm/structures"
)

type ListNode = structures.ListNode

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	hasCycle(structures.Ints2List([]int{1, 2, 3}))
}

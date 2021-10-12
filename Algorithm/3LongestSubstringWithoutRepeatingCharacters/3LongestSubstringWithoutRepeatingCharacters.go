package algorithm

func lengthOfLongestString(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]bool
	result, left, right := 0, 0, 0

	for i := 0; i <= len(s); i++ {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}

		if result < right-left {
			result = right - left
		}

		if left+result >= len(s) && right >= len(s) {
			break
		}
	}

	return result
}

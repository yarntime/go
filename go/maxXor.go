type TireNode struct {
	Next [2]*TireNode
}

func buildTree(nums []int) *TireNode {
	root := &TireNode {
		Next: [2]*TireNode{},
	}
	for _, num := range nums {
		cur := root
		for i := 31; i >= 0; i-- {
			pos := num & (1 << uint(i))
			if pos > 0 {
				pos = 1
			}
			if cur.Next[pos] == nil {
				cur.Next[pos] = &TireNode{
					Next: [2]*TireNode{},
				}
			}
			cur = cur.Next[pos]
		}
	}
	return root
}

func findMaxXor(num int, root *TireNode) int {
	result := 0
	for i := 31; i >= 0; i-- {
		pos := num & (1 << uint(i))
		if pos > 0 {
			pos = 1
		}
		if root.Next[1 - pos] != nil {
			result |= 1 << uint(i)
			root = root.Next[1 - pos]
		} else {
			root = root.Next[pos]
		}
	}
	return result
}

func findMaximumXOR(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	root := buildTree(nums)
	max := 0
	for _, num := range nums {
		r := findMaxXor(num, root)
		if r > max {
			max = r
		}
	}
	return max
}
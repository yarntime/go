func findKthLargest(nums []int, k int) int {
	heap := make([]int, k + 1)
	for i := 1; i <= k; i++ {
		c := i
		heap[c] = nums[i - 1]
		p := i / 2
		for p > 0 {
			if heap[p] <= heap[c] {
				break
			}
			heap[p], heap[c] = heap[c], heap[p]
            c = p
			p /= 2
		}
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < heap[1] {
			continue
		}
		p := 1
		heap[p] = nums[i]
		for p <= k / 2 {
			l, r := 2 * p, 2 * p + 1
			next := l
			if r <= k && heap[l] > heap[r] {
				next = r
			}
			if heap[p] <= heap[next] {
				break
			}
			heap[p], heap[next] = heap[next], heap[p]
			p = next
		}
	}

	return heap[1]
}
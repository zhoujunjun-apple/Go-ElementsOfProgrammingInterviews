package main

//ThreeBucket function: O(N) time and O(N) space
func ThreeBucket(nums []int, pivotIdx int) (newPivotIdx int) {
	newPivotIdx = pivotIdx
	if pivotIdx < 0 || pivotIdx >= len(nums) {
		return
	}

	less, equal, large := []int{}, []int{}, []int{}

	pivot := nums[pivotIdx]
	//distribute items to buckets
	for _, item := range nums {
		if item > pivot {
			large = append(large, item)
		} else if item < pivot {
			less = append(less, item)
		} else {
			equal = append(equal, item)
		}
	}

	//rearrange elements
	idx := 0
	for _, item := range less {
		nums[idx] = item
		idx++
	}
	newPivotIdx = idx
	for _, item := range equal {
		nums[idx] = item
		idx++
	}
	for _, item := range large {
		nums[idx] = item
		idx++
	}
	return
}

func main() {

}

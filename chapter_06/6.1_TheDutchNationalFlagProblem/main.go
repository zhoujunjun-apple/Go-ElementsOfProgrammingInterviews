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

//BruteForce function first rearrange less-than-pivot part, then rearrange large-than-pivot part.
//O(1) space and O(N*N) time
func BruteForce(nums []int, pivotIdx int) (newPivotIdx int) {
	newPivotIdx = pivotIdx
	if pivotIdx < 0 || pivotIdx >= len(nums) {
		return
	}

	pivotVal := nums[pivotIdx]

	lessTop := 0
	newPivotIdx = lessTop
	for ; lessTop < len(nums); lessTop++ {
		for lookIdx := lessTop + 1; lookIdx < len(nums); lookIdx++ {
			if nums[lookIdx] < pivotVal {
				swap(nums, lookIdx, lessTop)
				newPivotIdx = lessTop + 1
				break
			}
		}
	}
	//at this point, all items less than pivot has been put at right position

	for largeTop := len(nums) - 1; largeTop >= newPivotIdx; largeTop-- {
		for lookIdx := largeTop - 1; lookIdx >= newPivotIdx; lookIdx-- {
			if nums[lookIdx] > pivotVal {
				swap(nums, lookIdx, largeTop)
				break
			}
		}
	}
	//at that point, all items larger than pivot has been put at right positon
	//items equal to pivot automatically been put at the middle of array

	return
}

//DoubleScan function improve BruteForce solution to O(N) time
func DoubleScan(nums []int, pivotIdx int) (newPivotIdx int) {
	newPivotIdx = pivotIdx
	if pivotIdx < 0 || pivotIdx >= len(nums) {
		return
	}

	pivotVal := nums[pivotIdx]

	lessTop := 0
	newPivotIdx = lessTop
	for lookIdx := 0; lookIdx < len(nums); lookIdx++ {
		if nums[lookIdx] < pivotVal {
			swap(nums, lookIdx, lessTop)
			lessTop++
			newPivotIdx = lessTop
		}
	}

	largeTop := len(nums) - 1
	for lookIdx := len(nums) - 1; lookIdx >= newPivotIdx; lookIdx-- {
		if nums[lookIdx] > pivotVal {
			swap(nums, lookIdx, largeTop)
			largeTop--
		}
	}
	return
}

//SingleScan improve code simplisity by combining two scan from DoubleScan into single scan
//[0, less) contains items less than pivot;
//[less, equal) contains items equal than pivot;
//[equal, large) contains items to be classified;
//[large, len(nums)-1] contains items larger than pivot;
func SingleScan(nums []int, pivotIdx int) (newPivotIdx int) {
	newPivotIdx = pivotIdx
	if pivotIdx < 0 || pivotIdx >= len(nums) {
		return
	}

	pivotVal := nums[pivotIdx]
	less, equal, large := 0, 0, len(nums)
	for equal < large {
		item := nums[equal]
		if item > pivotVal {
			large-- //don't update equal, because swaped nums[large] need to recheck
			swap(nums, equal, large)
		} else if item < pivotVal {
			swap(nums, equal, less)
			equal++ //update equal, because nums[equal] have in right position, no need to recheck
			less++
		} else {
			equal++
		}
	}
	newPivotIdx = less
	return
}

//SingleScanBackward function like SingleScan, but use a backward scan order
//[0, less] contains items less than pivot;
//(less, equal] contains items not classified;
//(equal, large] contains items equal to pivot;
//(large, len(nums)-1] contains items large than pivot;
func SingleScanBackward(nums []int, pivotIdx int) (newPivotIdx int) {
	newPivotIdx = pivotIdx
	if pivotIdx < 0 || pivotIdx >= len(nums) {
		return
	}

	pivotVal := nums[pivotIdx]
	less, equal, large := -1, len(nums)-1, len(nums)-1
	for equal > less {
		item := nums[equal]
		if item > pivotVal {
			swap(nums, equal, large)
			large--
			equal--
		} else if item < pivotVal {
			less++
			swap(nums, less, equal)
		} else {
			equal--
		}
	}
	newPivotIdx = equal + 1
	return
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {

}

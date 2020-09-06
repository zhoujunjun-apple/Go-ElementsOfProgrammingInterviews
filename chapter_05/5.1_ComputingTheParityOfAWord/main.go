package main

//BruteForceParity function iteratively check each bit of word. O(N) time. N is the word bit length
func BruteForceParity(word int8) int8 {
	var oneCount int8 = 0
	var one int8 = 1
	for word > 0 {
		oneCount += word & one // "word & 1" extract the lowest bit from word
		word = word >> 1
	}
	return oneCount % 2
}

//BruteForceXORParity function iteratively check each bit of word with XOR operation. O(N) time
func BruteForceXORParity(word int8) int8 {
	var oneCount int8 = 0
	var one int8 = 1

	for word > 0 {
		oneCount ^= word & one //use XOR to count
		word = word >> 1
	}

	return oneCount
}

//DropLowestSetParity function drops the lowest set bit of word iteratively. Worst O(N) time
func DropLowestSetParity(word int8) int8 {
	var oneCount int8 = 0
	var one int8 = 1

	for word > 0 {
		oneCount ^= one
		word &= (word - 1) // this operation drops lowest set bit of word
	}

	return oneCount
}

func main() {}

package util

func BinarySearchInt(s []int, num, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if mid >= len(s) {
		return -1
	}

	if s[mid] == num {
		return mid
	} else if s[mid] > num {
		return BinarySearchInt(s, num, start, mid-1)
	}
	return BinarySearchInt(s, num, mid+1, end)
}

func FindIndex(s []int, val int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == val {
			return i
		}
	}
	return -1
}

func FindIndexString(s []string, val string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == val {
			return i
		}
	}
	return -1
}

func IsInSliceInt(array []int, val int) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

func IsInSliceUint8(array []uint8, val uint8) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

func IsIntArrayContainsSameValues(array []int) bool {
	if len(array) == 0 {
		return true
	}
	first := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] != first {
			return false
		}
	}
	return true
}

func NewArrayByteFilledWith(length int, defaultValue byte) []byte {
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		ret[i] = defaultValue
	}
	return ret
}

func NewIntArrayFilled(length int, from int) []int {
	ret := make([]int, length)
	for i := 0; i < length; i++ {
		ret[i] = from + i
	}
	return ret
}

func NewInt64ArrayFilledWith(length int, defaultValue int64) []int64 {
	ret := make([]int64, length)
	for i := 0; i < length; i++ {
		ret[i] = defaultValue
	}
	return ret
}

func NewByteArrayFilledWith(length int, defaultValue byte) []byte {
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		ret[i] = defaultValue
	}
	return ret
}

func CalcTotalArrayInt(a []int) int {
	var total = 0
	for _, v := range a {
		total += v
	}
	return total
}

func CalcTotalArrayInt64(a []int64) int64 {
	var total int64 = 0
	for _, v := range a {
		total += v
	}
	return total
}

func CalcTotalArrayFloat64(a []float64) float64 {
	var total float64 = 0
	for _, v := range a {
		total += v
	}
	return total
}

func CalcTotalLengthDoubleByteArray(array [][]byte) int {
	total := len(array)
	for _, e := range array {
		total += len(e)
	}
	return total
}

func ReverseSliceInt64(a []int64) []int64 {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func ReverseSliceInt(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

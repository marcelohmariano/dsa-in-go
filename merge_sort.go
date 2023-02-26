package godsa

// MergeSort is a divide-and-conquer sorting algorithm.
// Time complexity is O(nlog(n)).
// Space complexity is O(n).
func MergeSort(array []int) {
	buf := make([]int, 0, len(array))
	mergeSort(array, buf)
}

func mergeSort(array, buf []int) {
	if len(array) <= 1 {
		return
	}

	mid := len(array) / 2

	left := array[:mid]
	right := array[mid:]

	mergeSort(left, buf)
	mergeSort(right, buf)

	mergeArrays(array, buf, left, right)
}

func mergeArrays(array, buf []int, left, right []int) {
	var l, r int // left and right indexes

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			buf = append(buf, left[l])
			l++
		} else {
			buf = append(buf, right[r])
			r++
		}
	}

	buf = append(buf, left[l:]...)
	buf = append(buf, right[r:]...)

	copy(array, buf)
}

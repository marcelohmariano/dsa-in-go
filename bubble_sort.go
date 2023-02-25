package godsa

// BubbleSort is an in-place sorting algorithm.
// Time complexity is O(n^2).
// Space complexity is O(1).
func BubbleSort(array []int) {
	n := len(array)

	for n > 0 {
		newn := 0

		for i := 1; i < n; i++ {
			if array[i] < array[i-1] {
				array[i], array[i-1] = array[i-1], array[i]
				newn = i
			}
		}

		n = newn
	}
}

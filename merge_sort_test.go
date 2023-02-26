package godsa

import "testing"

func TestMergeSort(t *testing.T) {
	unordered := [...]int{3, 1, 5, 10, 2, 9, 8, 4, 6, 7}
	ordered := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	MergeSort(unordered[:])

	if unordered != ordered {
		t.Fatalf("MergeSort(3, 1, 5, 10, 2, 9, 8, 4, 6, 7) = %v, want %v", unordered, ordered)
	}
}

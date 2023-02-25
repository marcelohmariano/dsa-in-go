package godsa

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	unordered := [...]int{3, 1, 5, 10, 2, 9, 8, 4, 6, 7}
	ordered := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	BubbleSort(unordered[:])

	if unordered != ordered {
		t.Fatalf("BubbleSort(3, 1, 5, 10, 2, 9, 8, 4, 6, 7) = %v, want %v", unordered, ordered)
	}
}

package reverse

import "testing"

func TestReverseArray(t *testing.T) {
	expected := []int{3, 2, 1}
	result := reverseArray([]int{1, 2, 3})

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Reverse operation isn't right, got: %v, want: %v.", result, expected)
		}
	}
}

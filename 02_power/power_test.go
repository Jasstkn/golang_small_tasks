package main

import "testing"

func TestPawInt(t *testing.T) {
	expected := []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	result := pawInt(10)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Pow operation isn't right, got: %v, want: %v.", result, expected)
		}
	}
}

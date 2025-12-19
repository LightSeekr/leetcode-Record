package stack

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Example 1",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "Example 2",
			heights: []int{2, 4},
			want:    4,
		},
		{
			name:    "Single Element",
			heights: []int{1},
			want:    1,
		},
		{
			name:    "Empty Array",
			heights: []int{},
			want:    0,
		},
		{
			name:    "All Increasing",
			heights: []int{1, 2, 3, 4, 5},
			want:    9,
		},
		{
			name:    "All Decreasing",
			heights: []int{5, 4, 3, 2, 1},
			want:    9,
		},
		{
			name:    "Bug Case 1: Extension to Left",
			heights: []int{2, 1, 2},
			want:    3,
		},
		{
			name:    "Bug Case 2: Extension to Left",
			heights: []int{3, 2, 1},
			want:    4,
		},
		{
			name:    "All Same",
			heights: []int{2, 2, 2, 2},
			want:    8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("LargestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

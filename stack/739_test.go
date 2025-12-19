package stack

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{
			name:         "Example 1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "Example 2",
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			name:         "Example 3",
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
		{
			name:         "Decreasing temperatures",
			temperatures: []int{90, 80, 70, 60},
			expected:     []int{0, 0, 0, 0},
		},
		{
			name:         "Empty array",
			temperatures: []int{},
			expected:     []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("dailyTemperatures(%v) = %v; expected %v", tt.temperatures, result, tt.expected)
			}
		})
	}
}

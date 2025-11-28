package Hash_test

import (
	"reflect"
	"sort"
	"testing"

	"go_interview/leetcode/Hash"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 0},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{4, 2},
		},
		{
			name:   "with zero",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{3, 0},
		},
		{
			name:   "two elements",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{1, 0},
		},
		{
			name:   "large numbers",
			nums:   []int{1000000, 2000000, 3000000},
			target: 5000000,
			want:   []int{2, 1},
		},
		{
			name:   "no solution",
			nums:   []int{1, 2, 3},
			target: 10,
			want:   []int{0, 0},
		},
		{
			name:   "test2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{2, 1},
		},
		{
			name:   "test3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hash.TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// normalizeGroups sorts strings within each group, then sorts groups themselves.
// This allows comparison without depending on the order returned by GroupAnagrams.
func normalizeGroups(groups [][]string) [][]string {
	normalized := make([][]string, len(groups))
	for i, g := range groups {
		cp := append([]string(nil), g...)
		sort.Strings(cp)
		normalized[i] = cp
	}

	sort.Slice(normalized, func(i, j int) bool {
		gi, gj := normalized[i], normalized[j]
		if len(gi) != len(gj) {
			return len(gi) < len(gj)
		}
		for k := 0; k < len(gi) && k < len(gj); k++ {
			if gi[k] != gj[k] {
				return gi[k] < gj[k]
			}
		}
		return false
	})

	return normalized
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "basic example",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "single string",
			strs: []string{"abc"},
			want: [][]string{
				{"abc"},
			},
		},
		{
			name: "different length strings are not anagrams",
			strs: []string{"ab", "a"},
			want: [][]string{
				{"ab"},
				{"a"},
			},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hash.GroupAnagrams(tt.strs)
			gotNorm := normalizeGroups(got)
			wantNorm := normalizeGroups(tt.want)
			if !reflect.DeepEqual(gotNorm, wantNorm) {
				t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.strs, gotNorm, wantNorm)
			}
		})
	}
}

// go test -v -run TestLongestConsecutive
// go test -v -run TestLongestConsecutive
func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic example",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "all consecutive",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "no consecutive",
			nums: []int{1, 3, 5, 7, 9},
			want: 1,
		},
		{
			name: "with duplicates",
			nums: []int{1, 2, 0, 1},
			want: 3,
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, 0, 1, 2},
			want: 5,
		},
		{
			name: "large gap",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hash.LongestConsecutive(tt.nums)
			if got != tt.want {
				t.Errorf("LongestConsecutive(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

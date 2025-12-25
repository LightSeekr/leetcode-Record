package binarySearch

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			name: "case 1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 4,
			},
			want: []int{-1, -1},
		},
		{
			name: "case 2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "case 3",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "case 4",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{-1, -1},
		},
		{
			name: "case 5",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},
		{
			name: "case 6",
			args: args{
				nums:   []int{2, 2},
				target: 3,
			},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

package twoPointers

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := maxArea1(tt.args.height), maxArea(tt.args.height); got != want {
				t.Errorf("maxArea() = %v, want %v", got, want)
			}
		})
	}
}

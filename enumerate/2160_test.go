package enumerate

import "testing"

func TestMaximumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{9, 4, 3, 2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumDifference(tt.args.nums); got != tt.want {
				t.Errorf("MaximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

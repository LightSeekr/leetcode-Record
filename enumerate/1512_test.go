package enumerate

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
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
				nums: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 3, 1, 1, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIdenticalPairs(tt.args.nums); got != tt.want {
				t.Errorf("NumIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

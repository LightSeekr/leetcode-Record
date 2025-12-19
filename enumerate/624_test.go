package enumerate

import "testing"

func TestMaxDistance(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arrays: [][]int{
					{1, 2, 3},
					{4, 5},
					{1, 2, 3},
				},
			},
			want: 4,
		},

		{
			name: "case 2",
			args: args{
				arrays: [][]int{
					{1, 5},
					{3, 4},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDistance(tt.args.arrays); got != tt.want {
				t.Errorf("MaxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

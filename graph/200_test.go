package graph

import "testing"

func Test_numIslandsBfs(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := numIslandsBfs(tt.args.grid), numIslandsDfs(tt.args.grid); got != want {
				t.Errorf("numIslandsBfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

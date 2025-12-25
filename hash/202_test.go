package hash

import "testing"

func TestIsHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				n: 19,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHappy(tt.args.n); got != tt.want {
				t.Errorf("IsHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}

package stack

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				s: "([])",
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				s: "([)]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

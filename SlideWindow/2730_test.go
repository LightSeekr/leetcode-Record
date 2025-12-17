package SlideWindow

import "testing"

func TestLongestSemiRepetitiveSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "52233",
			},
			want: 4,
		},

		{
			name: "case 2",
			args: args{
				s: "1111111",
			},
			want: 2,
		},

		{
			name: "case 3",
			args: args{
				s: "5494",
			},
			want: 4,
		},
		{
			name: "case 4",
			args: args{
				s: "0001",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestSemiRepetitiveSubstring(tt.args.s); got != tt.want {
				t.Errorf("LongestSemiRepetitiveSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

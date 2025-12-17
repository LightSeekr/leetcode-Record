package SlideWindow

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
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
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "case 6",
			args: args{
				s: "aab",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

package strings

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
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
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				s: "   fly me   to   the moon  ",
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				s: "luffy is still joyboy",
			},
			want: 6,
		},
		{
			name: "case 4",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				s: "day",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

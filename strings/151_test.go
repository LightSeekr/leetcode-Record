package strings

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got2 := reverseWords(tt.args.s), reverseWords2(tt.args.s); got != tt.want || got2 != tt.want {
				t.Errorf("reverseWords() = %v,reverseWords2() = %v want %v", got, got2, tt.want)
			}
		})
	}
}

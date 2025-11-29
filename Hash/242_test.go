package Hash

import "testing"

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

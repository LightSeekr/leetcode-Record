package hash_test

import (
	"go_interview/leetcode/hash"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},

		{
			name: "case 2",
			args: args{
				ransomNote: "aa",
				magazine:   "bb",
			},
			want: false,
		},

		{
			name: "case 3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash.CanConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

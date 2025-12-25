package hash

import "testing"

func TestIsIsomorphic(t *testing.T) {
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
			name: "case 1",
			args: args{
				"egg",
				"add",
			},
			want: true,
		},
		{
			name: "case 1",
			args: args{
				"egg",
				"add",
			},
			want: true,
		},

		{
			name: "case 2",
			args: args{
				"foo",
				"bar",
			},
			want: false,
		},

		{
			name: "case 3",
			args: args{
				"paper",
				"title",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

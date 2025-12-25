package twoPointers

import "testing"

func Test_trap1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "case 1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			wantAns: 6,
		},

		{
			name: "case 2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			wantAns: 9,
		},
		{
			name: "case 3",
			args: args{
				height: []int{4, 2, 3},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := trap1(tt.args.height); gotAns != tt.wantAns {
				t.Errorf("trap1() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_trap2(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "case 1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			wantAns: 6,
		},

		{
			name: "case 2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			wantAns: 9,
		},
		{
			name: "case 3",
			args: args{
				height: []int{4, 2, 3},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := trap2(tt.args.height); gotAns != tt.wantAns {
				t.Errorf("trap2() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_trap3(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "case 1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			wantAns: 6,
		},

		{
			name: "case 2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			wantAns: 9,
		},
		{
			name: "case 3",
			args: args{
				height: []int{4, 2, 3},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := trap3(tt.args.height); gotAns != tt.wantAns {
				t.Errorf("trap3() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

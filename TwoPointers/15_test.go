package TwoPointers

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeTriplets(triplets [][]int) [][]int {
	out := make([][]int, 0, len(triplets))
	for _, t := range triplets {
		if len(t) != 3 {
			continue
		}
		cp := []int{t[0], t[1], t[2]}
		sort.Ints(cp)
		out = append(out, cp)
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i][0] != out[j][0] {
			return out[i][0] < out[j][0]
		}
		if out[i][1] != out[j][1] {
			return out[i][1] < out[j][1]
		}
		return out[i][2] < out[j][2]
	})

	uniq := out[:0]
	for _, t := range out {
		if len(uniq) == 0 ||
			uniq[len(uniq)-1][0] != t[0] ||
			uniq[len(uniq)-1][1] != t[1] ||
			uniq[len(uniq)-1][2] != t[2] {
			uniq = append(uniq, t)
		}
	}
	return uniq
}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeTriplets(threeSum(append([]int(nil), tt.args.nums...)))
			want := normalizeTriplets(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("threeSum() = %v, want %v", got, want)
			}
		})
	}
}

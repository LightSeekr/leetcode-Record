package twoPointers

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name           string
		args           args
		want           int
		wantNumsPrefix []int // 期望修改后的数组前want个元素
	}{
		{
			name:           "基本测试-含有重复元素",
			args:           args{nums: []int{1, 1, 2}}, // 经典示例
			want:           2,
			wantNumsPrefix: []int{1, 2},
		},
		{
			name:           "无重复元素",
			args:           args{nums: []int{1, 2, 3, 4}}, // 全部唯一元素
			want:           4,
			wantNumsPrefix: []int{1, 2, 3, 4},
		},

		{
			name:           "单元素数组",
			args:           args{nums: []int{5}}, // 边界情况
			want:           1,
			wantNumsPrefix: []int{5},
		},
		{
			name:           "所有元素相同",
			args:           args{nums: []int{1, 1}}, // 极端情况
			want:           1,
			wantNumsPrefix: []int{1},
		},
		{
			name:           "多个重复序列",
			args:           args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, // 多个重复组
			want:           5,
			wantNumsPrefix: []int{0, 1, 2, 3, 4},
		},
		{
			name:           "负整数测试",
			args:           args{nums: []int{-10, -10, -5, 0, 0, 5, 5}}, // 包含负数
			want:           4,
			wantNumsPrefix: []int{-10, -5, 0, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建数组副本以验证修改后的内容
			numsCopy := make([]int, len(tt.args.nums))
			copy(numsCopy, tt.args.nums)

			// 调用待测试函数
			got := RemoveDuplicates(numsCopy)

			// 验证返回的长度
			if got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}

			// 验证修改后的数组前got个元素是否符合预期
			if got > 0 && len(numsCopy) >= got && len(tt.wantNumsPrefix) >= got {
				for i := 0; i < got; i++ {
					if numsCopy[i] != tt.wantNumsPrefix[i] {
						t.Errorf("RemoveDuplicates() modified nums[%d] = %v, want %v", i, numsCopy[i], tt.wantNumsPrefix[i])
					}
				}
			}
		})
	}
}

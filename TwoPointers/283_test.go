package TwoPointers

import "testing"

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int // 期望移动后的结果
	}{
		{
			name: "基本测试-移动零到末尾",
			args: args{nums: []int{0, 1, 0, 3, 12}}, // 经典示例
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "无零元素",
			args: args{nums: []int{1, 2, 3, 4}}, // 没有零元素
			want: []int{1, 2, 3, 4},             // 数组应保持不变
		},
		{
			name: "全零元素",
			args: args{nums: []int{0, 0, 0, 0}}, // 所有元素都是零
			want: []int{0, 0, 0, 0},             // 数组应保持不变
		},
		{
			name: "空数组",
			args: args{nums: []int{}}, // 边界情况
			want: []int{},             // 空数组应保持不变
		},
		{
			name: "单元素数组-非零",
			args: args{nums: []int{5}}, // 单个非零元素
			want: []int{5},
		},
		{
			name: "单元素数组-零",
			args: args{nums: []int{0}}, // 单个零元素
			want: []int{0},
		},
		{
			name: "零在开头",
			args: args{nums: []int{0, 0, 1, 2}}, // 多个零在开头
			want: []int{1, 2, 0, 0},
		},
		{
			name: "零在结尾",
			args: args{nums: []int{1, 2, 0, 0}}, // 多个零在结尾
			want: []int{1, 2, 0, 0},             // 数组应保持不变
		},
		{
			name: "交错分布的零和非零",
			args: args{nums: []int{0, 1, 0, 2, 0, 3, 0, 4}}, // 零和非零元素交替出现
			want: []int{1, 2, 3, 4, 0, 0, 0, 0},
		},
		{
			name: "包含负数的数组",
			args: args{nums: []int{0, -1, 0, -2, 3, 0}}, // 包含负数
			want: []int{-1, -2, 3, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建数组副本以验证修改后的内容
			numsCopy := make([]int, len(tt.args.nums))
			copy(numsCopy, tt.args.nums)

			// 调用待测试函数
			MoveZeroes(numsCopy)

			// 验证数组长度是否保持不变
			if len(numsCopy) != len(tt.want) {
				t.Errorf("MoveZeroes() resulted in length %v, want length %v", len(numsCopy), len(tt.want))
				return
			}

			// 验证每个元素是否符合预期
			for i := range numsCopy {
				if numsCopy[i] != tt.want[i] {
					t.Errorf("MoveZeroes() at index %d: got %v, want %v", i, numsCopy[i], tt.want[i])
				}
			}

			// 额外验证：确保所有零都在非零元素之后
			foundZero := false
			for _, num := range numsCopy {
				if num == 0 {
					foundZero = true
				} else if foundZero {
					// 如果在找到零之后又发现非零元素，则测试失败
					t.Errorf("MoveZeroes() failed: non-zero element %v found after zero", num)
				}
			}
		})
	}
}

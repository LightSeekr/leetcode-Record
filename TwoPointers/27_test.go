package TwoPointers

import "testing"

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name           string
		args           args
		want           int
		wantNumsPrefix []int // 期望修改后的数组前want个元素（不包含val的元素）
	}{
		{
			name:           "基本测试-移除特定元素",
			args:           args{nums: []int{3, 2, 2, 3}, val: 3}, // 经典示例
			want:           2,
			wantNumsPrefix: []int{2, 2},
		},
		{
			name:           "全部匹配-所有元素都要移除",
			args:           args{nums: []int{5, 5, 5, 5}, val: 5}, // 所有元素都等于val
			want:           0,
			wantNumsPrefix: []int{},
		},
		{
			name:           "无匹配-不需移除任何元素",
			args:           args{nums: []int{1, 2, 3, 4}, val: 5}, // 没有元素等于val
			want:           4,
			wantNumsPrefix: []int{1, 2, 3, 4},
		},
		{
			name:           "空数组",
			args:           args{nums: []int{}, val: 1}, // 边界情况
			want:           0,
			wantNumsPrefix: []int{},
		},
		{
			name:           "单元素数组-等于val",
			args:           args{nums: []int{7}, val: 7}, // 单个元素等于val
			want:           0,
			wantNumsPrefix: []int{},
		},
		{
			name:           "单元素数组-不等于val",
			args:           args{nums: []int{7}, val: 8}, // 单个元素不等于val
			want:           1,
			wantNumsPrefix: []int{7},
		},
		{
			name:           "多个分散重复元素",
			args:           args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, // 分散的待删除元素
			want:           5,
			wantNumsPrefix: []int{0, 1, 3, 0, 4}, // 实际返回的前5个元素应该不包含2
		},
		{
			name:           "待删除元素在开头",
			args:           args{nums: []int{4, 4, 0, 1, 2}, val: 4}, // val在数组开头
			want:           3,
			wantNumsPrefix: []int{0, 1, 2},
		},
		{
			name:           "待删除元素在结尾",
			args:           args{nums: []int{1, 2, 3, 4, 4}, val: 4}, // val在数组结尾
			want:           3,
			wantNumsPrefix: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建数组副本以验证修改后的内容
			numsCopy := make([]int, len(tt.args.nums))
			copy(numsCopy, tt.args.nums)

			// 调用待测试函数
			got := RemoveElement(numsCopy, tt.args.val)

			// 验证返回的长度
			if got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}

			// 验证修改后的数组前got个元素是否符合预期（不包含val）
			if got > 0 && len(numsCopy) >= got && len(tt.wantNumsPrefix) >= got {
				for i := 0; i < got; i++ {
					if numsCopy[i] != tt.wantNumsPrefix[i] {
						t.Errorf("RemoveElement() modified nums[%d] = %v, want %v", i, numsCopy[i], tt.wantNumsPrefix[i])
					}
				}
			}

			// 额外验证：确保修改后的数组前got个元素中不包含val
			for i := 0; i < got && i < len(numsCopy); i++ {
				if numsCopy[i] == tt.args.val {
					t.Errorf("RemoveElement() failed: nums[%d] = %v still contains value to remove", i, numsCopy[i])
				}
			}
		})
	}
}

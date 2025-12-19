package stack

type MinStack struct {
	// 两个栈，一个按照大小，小的在上面，一个正常
	normal []int
	min    []int // 维护最小值
}

func Constructor() MinStack {
	return MinStack{
		normal: make([]int, 0),
		min:    make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.normal = append(this.normal, val)

	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else {
		if val <= this.min[len(this.min)-1] {
			this.min = append(this.min, val)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.normal) == 0 {
		return
	} else {
		num := this.normal[len(this.normal)-1]
		this.normal = this.normal[:len(this.normal)-1]
		if num == this.min[len(this.min)-1] {
			this.min = this.min[:len(this.min)-1]
		}
	}
}

func (this *MinStack) Top() int {
	//pop、top 和 getMin 操作总是在 非空栈 上调用
	if len(this.normal) == 0 {
		return 0
	} else {
		return this.normal[len(this.normal)-1]
	}
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	} else {
		return this.min[len(this.min)-1]
	}
}

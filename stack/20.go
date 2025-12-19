package stack

/*
*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
*/
type Stack[T string | int | rune | byte] struct {
	data []T
}

func NewStack[T string | int | rune | byte]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// Push 入栈
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {

	// 取最后一个元素
	idx := len(s.data) - 1
	val := s.data[idx]
	// 缩短切片
	s.data = s.data[:idx]
	return val
}

// Peek 返回栈顶元素但不弹出
func (s *Stack[T]) Peek() T {

	return s.data[len(s.data)-1]
}

// IsEmpty 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size 返回栈中元素个数
func (s *Stack[T]) Size() int {
	return len(s.data)
}

func isValid(s string) bool {
	stack := NewStack[rune]()
	for _, ch := range s {
		// 左括号入栈
		if ch == '(' || ch == '[' || ch == '{' {
			stack.Push(ch)
		} else {
			// 右括号
			if !stack.IsEmpty() {
				pop := stack.Pop()
				if ch == ')' && pop != '(' {
					return false
				}
				if ch == ']' && pop != '[' {
					return false
				}
				if ch == '}' && pop != '{' {
					return false
				}
			} else {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
func IsValid(s string) bool {
	return isValid(s)
}

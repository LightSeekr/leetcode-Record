package twoPointers

/*
*
[2,0,2,1,1,0]

[0,0,1,1,2,2]
*/
func sortColors(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}

	tail := 0
	head := n - 1
	for i := 0; i <= head; {
		if nums[i] == 0 {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail++
			i++
		} else if nums[i] == 2 {
			nums[head], nums[i] = nums[i], nums[head]
			head--
		} else {
			i++
		}
	}
}

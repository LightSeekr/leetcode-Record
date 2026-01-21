package design

import "math/rand"

type RandomizedSet struct {
	nums []int
	Map  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		Map:  make(map[int]int),
	}
}

//Insert(val):
/**
检查 Map，如果已存在，返回 false。
不存在：
把 val 加到数组末尾。
在 Map 中记录 val 的下标（就是数组当前的最后一位）。
返回 true。
*/
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Map[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.Map[val] = len(this.nums) - 1
	return true
}

/*
*Remove(val):

检查 Map，如果不存在，返回 false。
存在：
从 Map 中拿到 val 的下标，记为 idx。
拿到数组最后一个元素 lastVal。
关键步骤：把数组 idx 位置的元素替换为 lastVal。
更新 Map：把 lastVal 的下标更新为 idx。
删除数组最后一个元素（缩容）。
删除 Map 中的 val。
返回 true。
*/
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.Map[val]; !ok {
		return false
	} else {
		// 获取最后一个元素的值
		lastVal := this.nums[len(this.nums)-1]
		this.nums[idx] = lastVal
		this.nums = this.nums[:len(this.nums)-1]

		this.Map[lastVal] = idx
		delete(this.Map, val)
		return true
	}
}

/**
GetRandom():
生成一个 [0, len(nums)) 范围内的随机数 i。
返回 nums[i]。
*/

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

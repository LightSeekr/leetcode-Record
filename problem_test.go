package leetcode_test

import (
	"fmt"
	"go_interview/leetcode"
	"go_interview/leetcode/LinkedList"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	size := leetcode.RemoveElement(nums, val)
	fmt.Println(size)
	for i := 0; i < size; i++ {
		fmt.Println(nums[i])
	}
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6}
	res := leetcode.RemoveDuplicates(nums, 4)
	fmt.Println(res)
	for i := 0; i < res; i++ {
		fmt.Printf("%d ", nums[i])
	}

	fmt.Println("\n========================================")
	for i := res; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

func TestLRUCache(t *testing.T) {
	//["LRUCache","get","put","get","put","put","get","get"]
	//[   [2],     [2],  [2,6],[1], [1,5],[1,2], [1],  [2]]
	l := LinkedList.Constructor(2)
	fmt.Println("constructor LRU")
	fmt.Println("========================================")
	val := l.Get(2)
	fmt.Println("get 2:", val)

	l.Put(2, 6)

	val = l.Get(1)
	fmt.Println("get 1:", val)

	l.Put(1, 5)
	l.Put(1, 2)
	val = l.Get(1)
	fmt.Println("get 1:", val)

	val = l.Get(2)
	fmt.Println("get 2:", val)

	//l.Put(1, 0)
	//l.Put(2, 2)
	//fmt.Println("put [1,0] [2,2]")
	//printList(l.DList.Head)
	//printMap(l.M1)
	//fmt.Println("========================================")
	//
	//val := l.Get(1)
	//fmt.Println(val) //1
	//fmt.Println("get 1 return 0 ")
	//printList(l.DList.Head)
	//printMap(l.M1)
	//fmt.Println("========================================")
	//
	//l.Put(3, 3)
	//fmt.Println("put [3,3]")
	//printList(l.DList.Head)
	//printMap(l.M1)
	//fmt.Println("========================================")
	//
	//val = l.Get(2) // 返回 -1 (未找到)
	//fmt.Println(val)
	//fmt.Println("get 2 return -1 ")
	//printList(l.DList.Head)
	//printMap(l.M1)
	//fmt.Println("========================================")
	//
	//l.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	//fmt.Println("put [4,4]")
	//printList(l.DList.Head)
	//printMap(l.M1)
	//fmt.Println("========================================")
	//
	//val = l.Get(1) // 返回 -1 (未找到)
	//fmt.Println(val)
	//printMap(l.M1)
	//fmt.Println("\n========================================")
	//val = l.Get(3) // 返回 3
	//fmt.Println(val)
	//val = l.Get(4) // 返回 4
	//fmt.Println(val)
}

func printList(head *LinkedList.Node) {
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
}

func printMap(m map[int]*LinkedList.Node) {
	if len(m) == 0 {
		fmt.Println("map is empty")
		return
	}
	for k, v := range m {
		fmt.Printf("%d ", k)
		fmt.Printf("%d\n", v.Val)
	}

}

func TestCanConstruct(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"
	res := leetcode.CanConstruct(ransomNote, magazine)
	fmt.Println(res)
}

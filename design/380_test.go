package design

import (
	"fmt"
	"testing"
)

// ["RandomizedSet","remove","remove","insert","getRandom","remove","insert"]
// [[],[0],[0],[0],[],[0],[0]]
// [null,false,false,true,0,true,true]
func TestRandomizedSet(t *testing.T) {
	randomizedSet := Constructor()
	res := randomizedSet.Remove(0)
	fmt.Println("res=false", res)
	res = randomizedSet.Remove(0)
	fmt.Println("res=false", res)
	res = randomizedSet.Insert(0)
	fmt.Println("res=true", res)
	random := randomizedSet.GetRandom()
	fmt.Println("random=", random)
	res = randomizedSet.Remove(0)
	fmt.Println("res=true", res)
	res = randomizedSet.Insert(0)
	fmt.Println("res=true", res)
}

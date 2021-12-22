package offer

import "math/rand"

/*
https://leetcode-cn.com/problems/FortPu/
主要是通过map和list实现插入的O(1)
删除时通过只删最后一个实现删除的O(1)
随机访问通过list实现
list是动态数组
*/

type RandomizedSet struct {
	list []int
	mp   map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make([]int, 0), make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.mp[val]; ok {
		return false
	}
	r.mp[val] = len(r.list)
	r.list = append(r.list, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.mp[val]; !ok {
		return false
	}
	n := len(r.list)
	index := r.mp[val]
	r.mp[r.list[n-1]] = index
	delete(r.mp, val)
	r.list[index], r.list[n-1] = r.list[n-1], r.list[index]
	r.list = r.list[:n-1]
	return true
}

/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	n := len(r.list)
	return r.list[rand.Intn(n)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

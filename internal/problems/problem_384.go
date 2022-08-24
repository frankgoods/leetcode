package problems

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	cpy := append([]int{}, this.nums...)
	for i := range cpy {
		ii := rand.Intn(len(this.nums))
		cpy[i], cpy[ii] = cpy[ii], cpy[i]
	}
	return cpy
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

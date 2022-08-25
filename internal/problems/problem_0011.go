package problems

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	n := len(height)
	toCheck := []int{n - 1}
	for i := n - 2; i > 0; i-- {
		if height[i] > height[toCheck[len(toCheck)-1]] {
			toCheck = append(toCheck, i)
		}
	}

	fmt.Println(toCheck)

	m := 0
	left := []int{0}
	for i := 1; i < n; i++ {
		if height[i] > height[left[len(left)-1]] {
			left = append(left, i)
		}

		if i == toCheck[len(toCheck)-1] {
			for j := 0; j < len(left); j++ {
				sqr := (i - left[j]) * min(height[i], height[left[j]])
				m = max(m, sqr)
			}
			toCheck = toCheck[:len(toCheck)-1]
		}

	}
	return m
}

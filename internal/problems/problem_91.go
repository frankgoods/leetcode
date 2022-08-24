package problems

import "fmt"

func numDecodings(s string) int {
	s = "9" + s
	l := len(s)
	a := make([]int, l)
	a[0] = 1

	for i := 1; i < l; i++ {
		if s[i] == '0' {
			if s[i-1] > '2' || s[i-1] == '0' {
				return 0
			}
			a[i] = a[i-2]
		} else if s[i-1] == '0' || s[i-1] > '2' || (s[i-1] == '2' && s[i] > '6') {
			a[i] = a[i-1]
		} else {
			if a[i-1] == 1 {
				a[i] = 2
			} else {
				a[i] = a[i-1] + a[i-2]
			}
		}
	}

	fmt.Println(a)
	return a[l-1]
}

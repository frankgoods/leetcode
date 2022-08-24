package problems

func coinChange(coins []int, amount int) int {
	a := make([]int, amount+1)

	if amount == 0 {
		return 0
	}

	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] != 0 {
			for _, c := range coins {
				if i+c < len(a) && (a[i+c] == 0 || a[i+c] > a[i]+1) {
					a[i+c] = a[i] + 1
				}
			}
		}
	}

	if a[amount] == 0 {
		return -1
	} else {
		return a[amount]
	}
}

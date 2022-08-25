package problems

import "math/big"

func fromArray(a []int) *big.Int {
	mpl := big.NewInt(1)
	sum := big.NewInt(0)
	for i := len(a) - 1; i >= 0; i-- {
		aBig := big.NewInt(int64(a[i]))
		sum.Add(sum, aBig.Mul(aBig, mpl))
		mpl.Mul(mpl, big.NewInt(10))
	}
	return sum
}

func toArray(val *big.Int) (res []int) {
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	for ; val.Cmp(zero) == 1; val.Div(val, ten) {
		mod := new(big.Int).Mod(val, ten)
		res = append(res, int(mod.Int64()))
	}
	return reverse(res)
}

func reverse(a []int) (res []int) {
	res = make([]int, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}
	return res
}

func addToArrayForm(num []int, k int) []int {
	n := fromArray(num)
	return toArray(n.Add(n, big.NewInt(int64(k))))
}

package problems

func letterCombinations(digits string) []string {
	letters := map[int][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}

	var res []string

	if len(digits) == 0 {
		return res
	}

	comb(letters, digits, 0, "", &res)
	return res
}

func comb(letters map[int][]byte, dig string, i int, curCom string, res *[]string) {
	if i == len(dig) {
		*res = append(*res, curCom)
		return
	}

	curDig := int(dig[i] - '0')
	for _, l := range letters[curDig] {
		comb(letters, dig, i+1, curCom+string(l), res)
	}
}

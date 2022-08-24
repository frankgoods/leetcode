package problems

func gen(n, c int, cur string, res *[]string) {
	if n == 0 && c == 0 {
		*res = append(*res, cur)
		return
	}

	if n > 0 && c > 0 {
		gen(n-1, c+1, cur+"(", res)
		gen(n, c-1, cur+")", res)
	} else if n > 0 {
		gen(n-1, c+1, cur+"(", res)
	} else {
		gen(n, c-1, cur+")", res)
	}
}

func generateParenthesis(n int) []string {
	var res []string
	gen(n, 0, "", &res)
	return res
}

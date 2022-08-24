package problems

func canBreak(s string, pos int, wordDict, memo map[string]bool) bool {
	if pos == 0 {
		if v, ex := memo[s]; ex {
			return v
		}
		if _, exists := wordDict[s]; exists {
			return true
		}
	}
	if pos == len(s)-1 {
		return false
	}

	if _, exists := wordDict[s[0:pos+1]]; exists {
		v1 := canBreak(s, pos+1, wordDict, memo)
		v2 := canBreak(s[pos+1:], 0, wordDict, memo)
		memo[s[pos+1:]] = v2
		return v1 || v2
	}

	return canBreak(s, pos+1, wordDict, memo)
}

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	memo := make(map[string]bool)

	return canBreak(s, 0, dict, memo)
}

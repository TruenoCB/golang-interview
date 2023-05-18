package kmp

func GetNext(s string) []int {
	next := make([]int, len(s))
	pre := 0
	for suf := 1; suf < len(s); suf++ {
		for pre > 0 && s[suf] != s[pre] {
			pre = next[pre-1]
		}
		if s[pre] == s[suf] {
			pre++
		}
		next[suf] = pre
	}
	return next
}

func StrStr(s1, s2 string) []int {
	res := []int{}
	next := GetNext(s2)
	j := 0
	for i := 0; i < len(s1); i++ {
		for j > 0 && s1[i] != s2[j] {
			j = next[j-1]
		}
		if s1[i] == s2[j] {
			j++
			if j == len(s2) {
				j = 0
				res = append(res, i-len(s2)+1)
			}
		}
	}
	return res
}

func FirstStrStr(s1, s2 string) int {
	next := GetNext(s2)
	j := 0
	for i := 0; i < len(s1); i++ {
		for j > 0 && s1[i] != s2[j] {
			j = next[j-1]
		}
		if s1[i] == s2[j] {
			j++
			if j == len(s2) {
				return i - len(s2) + 1
			}
		}
	}
	return -1
}

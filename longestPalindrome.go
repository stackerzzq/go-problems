package main

//中心扩散
func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}

	rs := make([]byte, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		rs[2*i] = '#'
		rs[2*i+1] = []byte(s[i : i+1])[0]
	}
	rs[len(rs)-1] = '#'

	spread := func(rs []byte, center int) (step int) {
		for left, right := center-1, center+1; left >= 0 && right < len(rs) && rs[left] == rs[right]; left, right = left-1, right+1 {
			step++
		}

		return
	}

	var (
		maxLen, index int
	)

	for i := 0; i < len(rs); i++ {
		step := spread(rs, i)
		if step > maxLen {
			maxLen = step
			index = (i - maxLen) / 2
		}
	}

	return s[index : index+maxLen]
}

//manacher算法
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	rs := make([]byte, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		rs[2*i] = '#'
		rs[2*i+1] = []byte(s[i : i+1])[0]
	}
	rs[len(rs)-1] = '#'

	var maxRight, mirror, center, start, maxLen int
	p := make([]int, len(rs))

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for i := 0; i < len(rs); i++ {
		if i < maxRight {
			mirror = 2*center - i
			p[i] = min(maxRight-i, p[mirror])
		}

		for i-1-p[i] >= 0 && i+1+p[i] < len(rs) && rs[i-1-p[i]] == rs[i+1+p[i]] {
			p[i]++
		}

		if i+p[i] > maxRight {
			maxRight = i + p[i]
			center = i
		}

		if p[i] > maxLen {
			maxLen = p[i]
			start = (i - maxLen) / 2
		}
	}

	return s[start : start+maxLen]
}

// func main() {
//	ss := []string{"babad", "aaaaa"}
//	for _, s := range ss {
//		// println(longestPalindrome3(s))
//		println(longestPalindrome(s))
//	}
// }

package validanagram

func isAnagram(s string, t string) bool {
	m := [26]int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for i := 0; i < len(m); i++ {
		if m[i] != 0 {
			return false
		}
	}
	return true
}

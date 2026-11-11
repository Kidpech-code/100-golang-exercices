package practice

// Middle3 returns the 3 middle elements of s.
// If len(s) < 3, it returns s unchanged.
func Middle3(s []int) []int {
	if len(s) < 3 {
		return s
	}
	start := len(s)/2 - 1
	return s[start : start+3]
}

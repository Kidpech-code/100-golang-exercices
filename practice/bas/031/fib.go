package practice

func Fib(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}
	prev := Fib(n - 1)
	last := len(prev) - 1
	next := prev[last] + prev[last-1]
	return append(prev, next)
}

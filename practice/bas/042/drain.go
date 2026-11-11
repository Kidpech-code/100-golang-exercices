package practice

func Drain(nums []int) []int {
	ch := make(chan int, len(nums))

	for _, n := range nums {
		ch <- n
	}
	close(ch)

	result := make([]int, 0, len(nums))
	for n := range ch {
		result = append(result, n)
	}

	return result
}

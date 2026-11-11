package practice

func f1(c chan<- string, msg string) {
	c <- msg
}

func f2(c <-chan string) string {
	msg := <-c
	return "I am f2 and " + msg
}

func Relay(msg string) string {
	ch := make(chan string)
	go f1(ch, msg)
	return f2(ch)
}

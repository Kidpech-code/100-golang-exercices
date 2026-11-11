package practice

import (
	"fmt"
	"sync"
	"testing"
)

func TestRelayBas(t *testing.T) {
	got := Relay("Hello from f1")
	want := "I am f2 and Hello from f1"
	if got != want {
		t.Fatalf("Relay() = %q, want %q", got, want)
	}
}

func TestRelayBasConcurrent(t *testing.T) {
	const n = 10
	var wg sync.WaitGroup
	errCh := make(chan error, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			msg := fmt.Sprintf("msg-%d", i)
			want := "I am f2 and " + msg
			if got := Relay(msg); got != want {
				errCh <- fmt.Errorf("Relay(%q) = %q, want %q", msg, got, want)
			}
		}(i)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Error(err)
	}
}

package chapter5_singleton

import (
	"testing"
)

func TestGetOnlyOneEver(t *testing.T) {
	wait := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			_ = GetOnlyOneEver()

			// t.Run("", func(t *testing.T) {
			// })
		}
		wait <- 1
	}()
	<-wait
}

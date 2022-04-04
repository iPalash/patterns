package chapter5_singleton

import (
	"testing"
)

func TestGetOnlyOneEver(t *testing.T) {
	wait := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			t.Run("", func(t *testing.T) {
				g := GetOnlyOneEver()
				t.Log(g.id)
			})
		}
		wait <- 1
	}()
	<-wait
}

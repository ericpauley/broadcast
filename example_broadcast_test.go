package broadcast_test

import (
	"sync"
	"time"

	"github.com/ericpauley/broadcast"
)

func ExampleBroadcast() {
	var s sync.WaitGroup
	b := make(broadcast.Broadcast) // Broadcast channels are constructed exactly like regular channels
	s.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			b.Receive()
			time.Sleep(1 * time.Millisecond)
			s.Done()
		}()
	}
	b <- 1
	s.Wait()
}

package broadcast_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ericpauley/broadcast"
)

func BenchmarkNSubscribers1MessageBroadcast(bench *testing.B) {
	var s sync.WaitGroup
	s.Add(bench.N)
	b := make(broadcast.Broadcast)
	for i := 0; i < bench.N; i++ {
		go func() {
			b.Receive()
			s.Done()
		}()
	}
	time.Sleep(1 * time.Millisecond)
	b <- 10
	s.Wait()
}

func Benchmark1SubscriberNMessagesBroadcast(bench *testing.B) {
	var s sync.WaitGroup
	s.Add(1)
	b := make(broadcast.Broadcast)
	go func() {
		for i := 0; i < bench.N; i++ {
			b.Receive()
		}
		s.Done()
	}()
	time.Sleep(1 * time.Millisecond)
	for i := 0; i < bench.N; i++ {
		b <- 10
	}
	s.Wait()
}

func TestBroadcast(t *testing.T) {
	var s sync.WaitGroup
	b := make(broadcast.Broadcast)
	for i := 0; i < 10; i++ {
		s.Add(1)
		go func() {
			b.Receive()
			s.Done()
		}()
	}
	time.Sleep(1 * time.Millisecond)
	b <- 10
	s.Wait()
	close(b)
	b.Receive()
}

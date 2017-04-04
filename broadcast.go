// Package broadcast implements a channel which broadcasts messages to all listeners simultaneously.
package broadcast

type message struct {
	contents interface{}
	done     chan<- bool
}

// Broadcast represents a channel which can have a single item received by multiple goroutines.
//
// In every sense but receiving it can be used like a standard chan interface{}.
// Sends will block until at least one goroutine has received the interface.
// Sends will panic if the channel has been closed.
//
// Channels intended for broadcasting should NOT be listened to directly.
//
// If more than one channel is listening at the time the message is sent then they will all receive the message simultaneously.
type Broadcast chan interface{}

// Receive waits for a message to be received.
// If more receivers are listening then Receive will forward the packet to them as well.
func (b Broadcast) Receive() interface{} {
	defer recover() // Don't panic on rebroadcasts
	msg, ok := <-b
	if !ok {
		return nil
	}
	switch t := msg.(type) {
	case message:
		defer close(t.done)
		msg = t.contents
	}
	done := make(chan bool)
	select {
	case b <- message{msg, done}:
		<-done
	default:
	}
	return msg
}

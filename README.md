# Broadcast
Broadcast channels for Go! This is a lightweight implementation using just channels and no other synchronization

# How to Use:

Making a channel:

```Go
b := make(broadcast.Broadcast)
```

Broadcasting a message:

```Go
b <- message // Combine with whatever select magic you'd like!
```

Receiving a message:

```Go
message := b.Receive()
```

More info: https://godoc.org/github.com/ericpauley/broadcast

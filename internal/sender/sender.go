package sender

// Sender is an interface for sending txt messages
// and screenshots to receiver
type Sender interface {
	Send(msg string, photo []byte)
}

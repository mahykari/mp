package broker_sync

import "errors"

const BUFFER_SZ = 4

type BrokerSync struct {
	buffer chan string
}

func NewBrokerSync() *BrokerSync {
	return &BrokerSync{
		buffer: make(chan string, BUFFER_SZ),
	}
}

func (b *BrokerSync) Send(msg string) {
	b.buffer <- msg
}

func (b *BrokerSync) Receive() string {
	return <-b.buffer
}

func (b *BrokerSync) SendOrErr(msg string) error {
	select {
	case b.buffer <- msg:
		return nil
	default:
		return errors.New("buffer overflow")
	}
}

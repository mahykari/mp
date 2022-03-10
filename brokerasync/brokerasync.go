package brokerasync

import "errors"

const BUFFER_SZ = 4

type BrokerAsync struct {
	buffer chan string
}

func NewBrokerAsync() *BrokerAsync {
	return &BrokerAsync{
		buffer: make(chan string, BUFFER_SZ),
	}
}

func (b *BrokerAsync) Send(msg string) <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)
		b.buffer <- msg
		done <- true
	}()

	return done
}

func (b *BrokerAsync) Receive() <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		ch <- <-b.buffer
	}()

	return ch
}

func (b *BrokerAsync) SendOrErr(msg string) <-chan error {
	done := make(chan error)

	go func() {
		defer close(done)
		select {
		case b.buffer <- msg:
			done <- nil
		default:
			done <- errors.New("buffer overflow")
		}
	}()

	return done
}

func (b *BrokerAsync) ReceiveOrErr() (<-chan string, <-chan error) {
	ch := make(chan string)
	err := make(chan error)

	go func() {
		defer close(ch)
		defer close(err)
		select {
		case msg := <-b.buffer:
			ch <- msg
			err <- nil
		default:
			ch <- ""
			err <- errors.New("buffer underflow")
		}
	}()

	return ch, err
}

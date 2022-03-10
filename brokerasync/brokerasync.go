package brokerasync

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

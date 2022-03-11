package broker_multiclient

const BUFFER_SZ = 4
const N_CLIENTS = 4

type BrokerMultiClient struct {
	buffer []chan string
}

func NewBrokerMultiClient() *BrokerMultiClient {
	b := &BrokerMultiClient{
		buffer: make([]chan string, N_CLIENTS),
	}

	for i := 0; i < N_CLIENTS; i++ {
		b.buffer[i] = make(chan string, BUFFER_SZ)
	}

	return b
}

func (b *BrokerMultiClient) Send(i int, msg string) {
	b.buffer[i] <- msg
}

func (b *BrokerMultiClient) Receive(i int) string {
	return <-b.buffer[i]
}

package broker_multiclient

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestMultiClient(t *testing.T) {
	b := NewBrokerMultiClient()

	wg := &sync.WaitGroup{}

	client := func(i int) {
		defer wg.Done()
		msg := b.Receive(i % N_CLIENTS)
		log.Printf("CLIENT_%d: Recieving '%s'", i, msg)
	}

	for i := 0; i < N_CLIENTS; i++ {
		wg.Add(1)
		go client(i)
	}

	server := func() {
		defer wg.Done()
		for i := 0; i < N_CLIENTS; i++ {
			log.Printf("SERVER: Sending 'test%d'", i)
			b.Send(i%N_CLIENTS, fmt.Sprintf("test%d", i))
		}
	}

	wg.Add(1)
	go server()

	wg.Wait()
}

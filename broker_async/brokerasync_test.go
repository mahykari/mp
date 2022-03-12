package broker_async

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestAsync(t *testing.T) {
	b := NewBrokerAsync()
	ch := make([]<-chan bool, 2*BUFFER_SZ)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	server := func() {
		defer wg.Done()

		for i := 0; i < 2*BUFFER_SZ; i++ {
			ch[i] = b.Send("test" + fmt.Sprint(i))
			log.Println("sent", i)
		}

		for i := 0; i < 2*BUFFER_SZ; i++ {
			<-ch[i]
			log.Println("done", i)
		}
	}

	client := func() {
		defer wg.Done()

		time.Sleep(time.Second * 2)
		for i := 0; i < 2*BUFFER_SZ; i++ {
			log.Println("recieved", <-b.Receive())
		}
	}

	go server()
	go client()

	wg.Wait()
}

func TestBufferOverflow(t *testing.T) {
	b := NewBrokerAsync()
	ch := make([]<-chan error, BUFFER_SZ+1)

	for i := 0; i < BUFFER_SZ+1; i++ {
		ch[i] = b.SendOrErr("test" + fmt.Sprint(i))
	}

	for i := 0; i < BUFFER_SZ+1; i++ {
		log.Println("sent", i, "err", <-ch[i])
	}
}

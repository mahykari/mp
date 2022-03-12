package broker_sync

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	bs := NewBrokerSync()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	server := func() {
		defer wg.Done()
		for i := 0; i < 2*BUFFER_SZ; i++ {
			bs.Send("test" + fmt.Sprint(i))
			log.Println("sent", i)
		}
	}

	client := func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		for i := 0; i < 2*BUFFER_SZ; i++ {
			log.Println("recieved", bs.Receive())
		}
	}

	go server()
	go client()

	wg.Wait()
}

func TestBufferOverflow(t *testing.T) {
	bs := NewBrokerSync()
	for i := 0; i < BUFFER_SZ+1; i++ {
		err := bs.SendOrErr("test" + fmt.Sprint(i))
		log.Println("sent", i, "err", err)
	}
}

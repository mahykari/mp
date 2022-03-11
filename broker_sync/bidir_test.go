package broker_sync

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestBidir(t *testing.T) {
	bs := NewBrokerSync()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		bs.Send("test0")
		log.Printf("0: Sent 'test0'")
		time.Sleep(time.Millisecond * 200)
		log.Printf("0: Received '%s'", bs.Receive())
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 100)
		log.Printf("1: Received '%s'", bs.Receive())
		bs.Send("test1")
		log.Printf("1: Sent 'test1'")
	}()

	wg.Wait()
}

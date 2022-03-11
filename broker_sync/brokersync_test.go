package broker_sync

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSingleMessageSequential(t *testing.T) {
	bs := NewBrokerSync()
	bs.Send("test")
	if bs.Receive() != "test" {
		t.Error("Expected 'test'")
	}
}

func TestSingleMessageParallel(t *testing.T) {
	bs := NewBrokerSync()
	go bs.Send("test")
	if bs.Receive() != "test" {
		t.Error("Expected 'test'")
	}
}

func TestMultipleMessageSequential(t *testing.T) {
	bs := NewBrokerSync()
	messages := []string{"test1", "test2", "test3"}
	for _, msg := range messages {
		bs.Send(msg)
	}

	recieved := []string{}

	for i := 0; i < len(messages); i++ {
		recieved = append(recieved, bs.Receive())
	}

	for i := 0; i < len(messages); i++ {
		if messages[i] != recieved[i] {
			t.Error("Expected '" + messages[i] + "'")
		}
	}
}

func TestMultipleMessageParallel(t *testing.T) {
	bs := NewBrokerSync()
	messages := []string{"test1", "test2", "test3"}

	go func() {
		for _, msg := range messages {
			bs.Send(msg)
		}
	}()

	recieved := []string{}

	for i := 0; i < len(messages); i++ {
		recieved = append(recieved, bs.Receive())
	}

	for i := 0; i < len(messages); i++ {
		if messages[i] != recieved[i] {
			t.Error("Expected '" + messages[i] + "'")
		}
	}
}

func TestSync(t *testing.T) {
	bs := NewBrokerSync()
	go func() {
		for i := 0; i < 2*BUFFER_SZ; i++ {
			bs.Send("test" + fmt.Sprint(i))
			log.Println("sent", i)
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 2*BUFFER_SZ; i++ {
		log.Println("recieved", bs.Receive())
	}
}

func TestBufferOverflow(t *testing.T) {
	bs := NewBrokerSync()
	for i := 0; i < BUFFER_SZ+1; i++ {
		err := bs.SendOrErr("test" + fmt.Sprint(i))
		log.Println("sent", i, "err", err)
	}
}

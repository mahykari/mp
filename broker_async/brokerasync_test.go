package broker_async

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestSingleMessageSequential(t *testing.T) {
	b := NewBrokerAsync()
	b.Send("test")

	msg := <-b.Receive()

	if msg != "test" {
		t.Errorf("Expected 'test' but got '%s'", msg)
	}
}

func TestSingleMessageParallel(t *testing.T) {
	b := NewBrokerAsync()
	b.Send("test")

	done := make(chan struct{})

	go func() {
		msg := <-b.Receive()
		if msg != "test" {
			t.Errorf("Expected 'test' but got '%s'", msg)
		}
		done <- struct{}{}
	}()

	<-done
}

func TestMultipleMessageSequential(t *testing.T) {
	b := NewBrokerAsync()
	messages := []string{"test1", "test2", "test3"}

	for _, msg := range messages {
		b.Send(msg)
	}

	for i := 0; i < len(messages); i++ {
		msg := <-b.Receive()
		if !contains(messages, msg) {
			t.Errorf("Expected '%s' but got '%s'", messages[i], msg)
		}
	}
}

func TestMultipleMessageParallel(t *testing.T) {
	b := NewBrokerAsync()
	messages := []string{"test1", "test2", "test3"}

	for _, msg := range messages {
		go func(x string) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			b.Send(x)
		}(msg)
	}

	for i := 0; i < len(messages); i++ {
		msg := <-b.Receive()
		if !contains(messages, msg) {
			t.Errorf("Expected '%s' but got '%s'", messages[i], msg)
		}
	}
}

func TestAsync(t *testing.T) {
	done := make(chan struct{})
	b := NewBrokerAsync()
	ch := make([]<-chan bool, 2*BUFFER_SZ)

	go func() {
		for i := 0; i < 2*BUFFER_SZ; i++ {
			ch[i] = b.Send("test" + fmt.Sprint(i))
			log.Println("sent", i)
		}
		for i := 0; i < 2*BUFFER_SZ; i++ {
			<-ch[i]
			log.Println("done", i)
		}
		done <- struct{}{}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 2*BUFFER_SZ; i++ {
		log.Println("recieved", <-b.Receive())
	}

	<-done
}

func TestBufferOverflow(t *testing.T) {
	b := NewBrokerAsync()
	for i := 0; i < BUFFER_SZ+1; i++ {
		ch := b.SendOrErr("test" + fmt.Sprint(i))
		log.Println("sent", i, "err", <-ch)
	}
}

package brokerasync

import (
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
	ch := b.Send("test")

	done := <-ch
	if !done {
		t.Errorf("Expected true but got false")
	}

	msg := <-b.Receive()

	if msg != "test" {
		t.Errorf("Expected 'test' but got '%s'", msg)
	}
}

func TestSingleMessageParallel(t *testing.T) {
	b := NewBrokerAsync()
	ch := b.Send("test")

	go func() {
		done := <-ch

		if !done {
			t.Errorf("Expected true but got false")
		}

		msg := <-b.Receive()
		if msg != "test" {
			t.Errorf("Expected 'test' but got '%s'", msg)
		}
	}()
}

func TestMultipleMessageSequential(t *testing.T) {
	b := NewBrokerAsync()
	messages := []string{"test1", "test2", "test3"}

	for _, msg := range messages {
		ch := b.Send(msg)
		done := <-ch
		if !done {
			t.Errorf("Expected true but got false")
		}
	}

	for i := 0; i < len(messages); i++ {
		msg := <-b.Receive()
		if msg != messages[i] {
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
			ch := b.Send(x)
			done := <-ch
			if !done {
				t.Errorf("Expected true but got false")
			}
		}(msg)
	}

	for i := 0; i < len(messages); i++ {
		msg := <-b.Receive()
		if !contains(messages, msg) {
			t.Errorf("Expected '%s' but got '%s'", messages[i], msg)
		}
	}
}

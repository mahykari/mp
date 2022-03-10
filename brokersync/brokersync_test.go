package brokersync

import (
	"testing"
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

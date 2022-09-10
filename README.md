# MP

Imeplementation of a message broker with Go.

## Requirements

Make sure you have the Go programming language installed on your machine.
For simplicity, you might use Go in a set-up IDE, such as Visual Studio Code or GoLand.

## What to Run

For each type of broker (e.g. `broker_sync`), there exists at least a test file (`brokersync_test.go`).
You can either run or play around with these tests; they capture essential behavior of brokers.

## Features

This tiny library provides a sync, async, and (fixed-size) multi-client message broker.
Each broker has at least two methods: `Send` and `Receive`;
the sync and async brokers support a dedicated method `SendOrErr` to handle with buffer overflow.

## Sample Source

Each broker will have a structure similar to the following:

```go
type BrokerSample struct {
    buffer chan string
}

func NewBrokerSample() *BrokerSync {
    return &BrokerSync{
        buffer: make(chan string, BUFFER_SZ),
    }
}

func (b *BrokerSample) Send(msg string) {
    // ...
}

func (b *BrokerSample) Receive() string {
    // ...
}

// Other methods
```

## Notes

This project was meant for an assignment for the Distributed Systems course at University of Tehran, Spring 2022.

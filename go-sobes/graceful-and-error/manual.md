### Graceful and handling errors

```go
package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    if err := Main(); err != nil {
        log.Fatal(err)
    }
}

func Main() error {
    q, err := NewSendQueue()
    if err != nil {
        return fmt.Errorf("error creating send queue: %w", err)
    }
    defer func() {
        if err := q.Close(); err != nil {
            log.Printf("error stopping send queue: %v", err)
        }
    }()

    s, err := StartService(q)
    if err != nil {
        return fmt.Errorf("error starting service: %w", err)
    }
    defer func() {
        if err := s.Stop(); err != nil {
            log.Printf("error stopping service: %v", err)
        }
    }()

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    <-sigs // waiting for shutdown signal

    return nil
}

type Service struct{}

func StartService(*SendQueue) (*Service, error) {
    return &Service{}, nil
}

func (*Service) Stop() error {
    return errors.New("some error")
}

type SendQueue struct{}

func NewSendQueue() (*SendQueue, error) {
    return &SendQueue{}, nil
}

func (*SendQueue) Close() error {
    return errors.New("some error")
}

```

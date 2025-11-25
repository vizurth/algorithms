### Sleep implement
```go
package main
import (
    "context"
    "log"
    "time"
)
// Sleep is interruptible verion of time.Sleep.
// Returns false if context was cancelled,
// 1 способ
func Sleep(ctx context.Context, duration time.Duration) bool {
    timer := time.NewTimer(duration)
    defer timer.Close()
    
    select {
        case <-timer.C:
            return true
        case <-ctx.Done():
            return false
    }
}
// 2 способ
func Sleep(ctx context.Context, duration time.Duration) bool {
    select {
        case <-time.After(duration):
            return true
        case <-ctx.Done():
            return false
    }
}
func main() {
    ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)
    if !Sleep(ctx, time.Second) {
        log.Print("interrupted")
    }
}
‘‘
```

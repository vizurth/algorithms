### Merge Channels(Fan in)
```go
package main
func merge(chs ...chan int) chan int {
    out := make(chan int)
    
    wg := &sync.WaitGroup{}
    
    wg.Add(len(chs))
    for _, ch := range chs{
        go func() {
            defer wg.Done()
            for item := range ch {
               out <- item
            }
        }()
    }
    
    go func () {
       wg.Wait()
       close(out) 
    }()
    
    return out
}

func main() {
    ch1 := startProducerA()
    ch2 := startProducerB()
    for el := range merge(ch1, ch2) {
        println(el)
    }
}
```


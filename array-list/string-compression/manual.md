### 443. String Compression
```go
func compress(chars []byte) int {
    write, read := 0,0
    
    for read < len(chars) {
        char := chars[read]
        count := 0
        // читаем символы
        for read < len(chars) && chars[read] == char {
            read++
            count++
        }
        // записываем символ
        chars[write] = char
        write++

        // записываем число
        if count > 1{
            for _, digit := range strconv.Itoa(count) {
                chars[write] = byte(digit)
                write++
            }
        }
    }



    return write
}
```
### Описание алгоритма
Алгоритм использует два указателя:
- read — для чтения исходного массива;
- write — для записи результата.
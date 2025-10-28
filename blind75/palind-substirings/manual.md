### 7. Palindromic Substrings
```go
func countSubstrings(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        count += expandFromCenter(s, i, i)
        count += expandFromCenter(s, i, i+1)
    }
    return count
}

func expandFromCenter(s string, left, right int) int {
    count := 0
    for left >= 0 && right < len(s) && s[left] == s[right] {
        count++
        left--
        right++
    }
    return count
}
```
### Описание алгоритма
Идея заключается в том чтобы считать сколько встречается палиндромов расширяясь с какого-то индекса будем считаь для четных и не четных палиндромов
### 3. Longest Substring Without Repeating Characters
```go
//  sliding window
func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]bool)
    lenght := 0
    l, r := 0,0

    for r < len(s) {
        if _, ok := set[s[r]]; !ok {
            set[s[r]] = true
            lenght = max(lenght, r - l + 1)
            r++ 
        } else {
            delete(set, s[l])
            l++
        }
    }


    return lenght
}
```	
### Описание алгоритма
Идея алгоритма заключается в sliding window(два указателя)
будем расширять наше окно если увидели элемент которого нет в set и уменьшать его если элемент уже был в set

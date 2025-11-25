### 136. Single Number
```go
func singleNumber(nums []int) int {
    mask := 0
    for _, num := range nums {
        mask ^= num
    }

    return mask
}
```
### Описание алгоритма
Просто проходимся циклом по массиву и используем **XOR** если числа равны то на следующей итерации будет 0 иначе следующее число
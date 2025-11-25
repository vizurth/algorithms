### 11. Missing Number

```go
func missingNumber(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    lenght := len(nums)
    curr := ((1 + lenght)*lenght) / 2

    return curr - sum
}
```

### Описание алгоритма

Нужно посчитать корректную сумму которая должна быть как сумма ариф прогресси и вычесть сумма которая получилась в `nums`

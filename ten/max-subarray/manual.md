### 53. Maximum Subarray
```go
func maxSubArray(nums []int) int {
    max_sum := nums[0]
    curr_sum := nums[0]

    for i := 1; i < len(nums); i++{
        num := nums[i]
        curr_sum = max(curr_sum + num, num)
        max_sum = max(curr_sum, max_sum)
    }

    return max_sum
}
```
### Описание алгоритма
Ну будем фиксировать максимальную сумму и изменять curr_sum если при + num ничего хорошего не происходит

### 448. Find All Numbers Disappeared in an Array
```go
func findDisappearedNumbers(nums []int) []int {
    i := 0
    for i < len(nums) {
        pos := nums[i] - 1
        if nums[i] != nums[pos]{
            nums[i], nums[pos] = nums[pos], nums[i]
        } else {
            i++
        }
    }
    res := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] != i + 1{
            res = append(res, i + 1)
        }
    }

    return res
}
```
### Описание алгоритма
Используем Cyclic Sort, что будем делать будем менять местами цисла на свои позиции по индексам до того пока не дойдем до последнего элемента, а далее посмотрим стоят ли числа на своих местах.
### О большое
Время: O(n)
Память: O(n)
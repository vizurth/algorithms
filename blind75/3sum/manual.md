### 12. 3Sum
```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}

    for i:=0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left, right := i + 1, len(nums) - 1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            } else {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                // избавляемся от дубликатов
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            }
        }
    }
    return res
}
```
### Описание алгоритма
Сортируем массив. Будем проходиться по массиву `len(nums) - 2` так как закладываем место под left right. ставим `left = i + 1, right = len(nums) - 1`. зафиксировалли одно число и ищем для него `left` и `right` пропуская дубликаты для того чтобы посмотреть что там дальше
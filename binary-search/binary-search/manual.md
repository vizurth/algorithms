### 704. Binary Search
```go
func check(index int, nums []int, target int) bool {
    return nums[index] >= target
}

func leftBinSearch(l, r int, nums []int, target int) int {
    for l < r { // пока указатели не равны
        m := (l + r) / 2
        if check(m, nums, target) {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func search(nums []int, target int) int {
    index := leftBinSearch(0, len(nums)-1, nums, target)
    if nums[index] == target {
        return index
    } else {
        return -1
    }
}
```
### Описание алгоритма
использовать левый бинарный поиск

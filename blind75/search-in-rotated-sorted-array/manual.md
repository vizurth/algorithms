### 22. Search in Rotated Sorted Array
```go
func findPivotIdx(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        m := (l + r) / 2
        if nums[m] < nums[r] {
            r = m 
        } else {
            l = m + 1
        }
    }
    return l
}


func leftSearch(l, r, target int, nums []int) int {
    for l < r {
        m := (l + r) / 2
        if nums[m] >= target {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func search(nums []int, target int) int {
    pivotIdx := findPivotIdx(nums)
    left := leftSearch(0, pivotIdx-1, target, nums)
    right := leftSearch(pivotIdx, len(nums)-1, target, nums)

    if nums[left] == target {
        return left
    } else if nums[right] == target {
        return right
    } else {
        return -1
    }
}
```
### Описание алгоритма

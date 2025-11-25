### 33. Search in Rotated Sorted Array

#### Первый способ:
```go
func search(nums []int, target int) int {
    l, r := 0,  len(nums) - 1
    for l <= r{
        m := (l + r) / 2
        if nums[m] == target{
            return m
        }

        if nums[l] <= nums[m]{
            if nums[m] > target && nums[l] <= target{
                r = m - 1
            } else {
                l = m + 1
            }
        } else {
            if nums[m] < target && nums[r] >= target{
                l = m + 1
            } else {
                r = m - 1
            }
        }
        fmt.Println(l, m, r)
    }

    return -1
}
```
#### Второй способ:
```go
func pivot(nums []int) int {
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
    pivotIdx := pivot(nums)
    l := leftSearch(0, pivotIdx - 1, target, nums)
    r := leftSearch(pivotIdx, len(nums)-1, target, nums)
    fmt.Println(pivotIdx, l, r)
    if nums[l] == target {
        return l
    } else if nums[r] == target {
        return r
    }
    return -1
}
```
### Описание алгоритма
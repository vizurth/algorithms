### 410. Split Array Largest Sum
```go
func check(m, k int, nums []int) bool {
    count := 1
    sum := 0
    for _, num := range nums {
        sum += num
        if sum > m{
            count++
            sum = num
        }

    }
    return count <= k
}


func splitArray(nums []int, k int) int {
    l, r := 0, 0
    for _, num := range nums {
        r += num
        l = max(l, num)
    }

    for l < r {
        m := (l + r) / 2
        if check(m, k, nums) {
            r = m 
        } else {
            l = m + 1
        }
    }

    return l
}
```
### Описание алгоритма
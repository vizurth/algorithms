### 189. Rotate Array
```go
func reverse(a []int, start, end int){
    for start < end {
        a[start], a[end] = a[end], a[start]
        start++
        end--
    }
}
func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}
```
### Описание алгоритма
### 189. Rotate Array
```go
func reverse(a []int, start, end int){
    for start < end {
        a[start], a[end] = a[end], a[start]
        start++
        end--
    }
}
func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}
```

### Описание алгоритма

### 724. Find Pivot Index
```go
func prefSumm(nums []int) []int {
    pref := make([]int, len(nums) + 1)
    for i := 1; i< len(nums) + 1; i++ {
        pref[i] = pref[i-1] + nums[i-1]
    }

    return pref
} 

func rangeFunc(left, right int, pref *[]int) int {
    return (*pref)[right] - (*pref)[left]
}

func pivotIndex(nums []int) int {
    pref := prefSumm(nums)
    for i:=0;i<len(nums);i++{
        l := rangeFunc(0, i, &pref)
        r := rangeFunc(i+1, len(nums), &pref)
        if l == r {
            return i
        }
    }
    return -1
}
```
### Описание алгоритма 
посчитаем преф суммы с [0, i), [i+1, n) и сравним их
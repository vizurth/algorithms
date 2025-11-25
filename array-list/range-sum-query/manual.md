### 303. Range Sum Query - Immutable
```go
type NumArray struct {
    slice []int
    prefSlice []int
}
// префсуммы

func (n *NumArray) countPref(left, right int) int {
    return n.prefSlice[right+1] - n.prefSlice[left]
}

func Constructor(nums []int) NumArray {
    prefSlice := make([]int, len(nums) + 1)
    for i := 1; i< len(nums) + 1; i++ {
        prefSlice[i] = prefSlice[i-1] + nums[i-1]
    }
    return NumArray{
        slice: nums,
        prefSlice: prefSlice,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.countPref(left, right)
}
```
### Описание алгоритма
задачка на обычную преф сумму главное помнить что считается промедуто как **[left, right)** не включительно и все можно спокойно реши
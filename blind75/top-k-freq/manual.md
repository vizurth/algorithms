###

```go
func topKFrequent(nums []int, k int) []int {
    result := []int{}
    dict := make(map[int]int)
    for _, num := range nums{
        dict[num]++
    }
    slicefreq := make([][]int, 0)
    for key, value := range dict{
        slicefreq = append(slicefreq, []int{key, value})
    }
    sort.Slice(slicefreq, func(i, j int) bool {
        return slicefreq[i][1] > slicefreq[j][1]
    })
    for i := 0; i<k; i++ {
        result = append(result, slicefreq[i][0])
    }
    return result
}
```

### Описание алгоритма

проиводим подсчет всех элементов в массиве, далее делаем пары чисел сортируем по частоте и выводим первые k элементов

### О большое

Время: `O(n logn)`
Память: `O(n)` - когда все элементы уникальны

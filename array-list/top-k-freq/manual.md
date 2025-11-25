### 347. Top K Frequent Elements
```go
func topKFrequent(nums []int, k int) []int {
    dict := make(map[int]int)
    maxFreq := 0

    for _, num := range nums {
        dict[num]++
        if dict[num] > maxFreq {
            maxFreq = dict[num]
        }
    }

    sliceFreq := make([][]int, maxFreq+1)
    for num, freq := range dict {
        sliceFreq[freq] = append(sliceFreq[freq], num)
    }

    result := make([]int, 0, k)
    for i := maxFreq; i >= 1 && k > 0; i-- {
        for _, num := range sliceFreq[i] {
            result = append(result, num)
            k--
            if k == 0 {
                break
            }
        }
    }

    return result
}
```
### Описание алгоритма
Будет считать частоту в мапу и создадим массив где храняться частоты
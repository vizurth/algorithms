### 1. Longest Consecutive Sequence
```go
func longestConsecutive(nums []int) int {
    set := make(map[int]bool, 0)
    for _, num := range nums {
        set[num] = true
    }
    longest := 0

    for num := range set {
        if !set[num-1] {
            count := 0
            for set[num + count] {
                count++
            }
            longest = max(longest, count)
        }
    }

    return longest
}
```
### Описание алгоритма
Задача решается через hashmap, сначала мы заполняем set значениями из nums
после проходимся по set и смотрим если у выбранного нами числа нет прерыдущего тоесть num-1 можем построить последовательность
создаем переменную count которая будет считать сколько элементов в последовательности в начале count = 0 после уже через max будем считать наибольший элемент
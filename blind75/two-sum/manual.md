### 2. Two Sum
```go
func twoSum(nums []int, target int) []int {
    hashTable := make(map[int]int, 0)

    for i, elem := range nums {
        left := target - elem
        if idx, ok := hashTable[left]; ok {
            return []int{i, idx}
        }
        hashTable[elem] = i
    }
    return []int{}
}
```
### Описание алгоритма
Так как задача называется сумма двух то возвращать мы будем два элемента.
Задача также на hashmap в это раз будем хранить (число: его индекс в массиве) проходимся по nums и ищем есть ли в массиве число обратное ему(target - num) если есть возвращаем индекс (num, target-num) порядок не важен если числа нет записываем его в hashmap
### 217. Contains Duplicate
```go
func containsDuplicate(nums []int) bool {
    set := make(map[int]struct{}, 0)
    for _, elem := range nums {
		if _, ok := set[elem]; ok {
			return true
		}
        set[elem] = struct{}{}
    }

    return false
}
```
### Описание алгоритма
делаем cет и потом просто сравниваем встречали ли раньше или нет
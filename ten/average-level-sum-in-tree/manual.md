### 637. Average of Levels in Binary Tree

```go
func treeFunc(root *TreeNode, dict map[int][]int, level int) {
    if root == nil {
        return
    }
    dict[level] = append(dict[level], root.Val)
    treeFunc(root.Left, dict, level+1)
    treeFunc(root.Right, dict, level+1)
}

func averageOfLevels(root *TreeNode) []float64 {
    dict := make(map[int][]int)
    treeFunc(root, dict, 0)

    result := make([]float64, len(dict))
    for i := 0; i<(len(dict)); i++ {
        sum := 0
        for _, num  := range dict[i]{
            sum += num
        }
        result[i] = float64(sum)/float64(len(dict[i]))
    }

    return result
}
```

### Описание алгоритма

Будем работать с уровнями в дереве, и записывать элементы дерева в `map` после того как прошлись по всему дереву создаем пустой массив проходимся по его индексам и записываем в них среднее арифметическое элементов на каждом уровне

### О большое

Время: `O(N)`
Память: `O(N)`

### 111. Minimum Depth of Binary Tree
```go
func treeFunc(root *TreeNode, res *int, level int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        if level < *res {
            *res = level + 1
        }
        return
    }
    treeFunc(root.Left, res, level + 1)
    treeFunc(root.Right, res, level + 1)
}

func minDepth(root *TreeNode) int {
    minDepth := math.MaxInt
    treeFunc(root, &minDepth, 0)
    if minDepth == math.MaxInt {
        return 0
    }
    return minDepth
}
```

#### Описание алгоритма 
Проходимся по нашему дереву учитываем уровень если дошли до ноды у которой нет детей то смотрим на наш результат и записываем его если он меньше предыдущего

#### О большое
Время: `O(N)`
Память: `O(log N | H)` H - высота дерева
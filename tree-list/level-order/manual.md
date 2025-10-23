### 102. Binary Tree Level Order Traversal
```go
func prepareLevelOrder(root *TreeNode, result *[][]int, level int) {
    if root == nil {
        return
    }
    if len(*result) < level + 1 {
        *result = append(*result, []int{})
    }
    prepareLevelOrder(root.Left, result, level + 1)
    (*result)[level] = append((*result)[level], root.Val)
    prepareLevelOrder(root.Right, result, level + 1)
}


func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    prepareLevelOrder(root, &result, 0)
    return result
}
```
### Описание алгоритма
Просто храним левел и не забываем проверять на добавляемый массив в матрицу

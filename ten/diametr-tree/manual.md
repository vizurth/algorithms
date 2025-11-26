### 543. Diameter of Binary Tree
```go
func check(root *TreeNode, dia *int) int {
    if root == nil {
        return 0
    }
    left := check(root.Left, dia)
    right := check(root.Right, dia)
    *dia = max(*dia, left + right)
    return max(left, right) + 1
}
func diameterOfBinaryTree(root *TreeNode) int {
    dia := 0
    check(root, &dia)
    return dia
}
```

### Описание алгоритма
Так будем проходится по дереву в ширину и будем считать наш диаметр когда дошли до конца получает что left right == 0 следовательно left == 1 и тд до того момента пока не дойдем до вершины и не запишем диаметр наверх
### О большое
Время: `O(N)`
Память: `O(H)`
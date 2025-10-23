### 101. Symmetric Tree
```go 
func checkTree(l, r *TreeNode) bool {
    if l == nil || r == nil {
        return l == nil && r == nil 
    }
    if l.Val != r.Val {
        return false
    }
    return checkTree(l.Left, r.Right) && checkTree(l.Right, r.Left)
}


func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return checkTree(root.Left, root.Right)
}
```
### Описание алгоритма
Дерево считается симметричным, если:

левая и правая его поддеревья являются зеркальными копиями друг друга.

Чтобы это проверить, используется рекурсивная функция checkTree(l, r), которая сравнивает два поддерева:
	левое поддерево (l)
	и правое поддерево (r).

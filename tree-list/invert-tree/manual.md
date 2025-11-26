### 226. Invert Binary Tree

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    right := invertTree(root.Right)
    left := invertTree(root.Left)
    root.Left = right
    root.Right = left

    return root
}
```

#### Описание алгоритма

Просто рекурсивно проходимся и меняем местами указатели и все

#### О большое

Время: `O(N)`
Память: `O(H)`

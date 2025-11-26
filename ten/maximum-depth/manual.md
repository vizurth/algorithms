### 104. Maximum Depth of Binary Tree

```go
func check(root *TreeNode, res *int, lvl int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        if lvl + 1 > *res {
            *res = lvl + 1
        }
    }
    check(root.Left, res, lvl+1)
    check(root.Right, res, lvl+1)
}
func maxDepth(root *TreeNode) int {
    depth := math.MinInt
    check(root, &depth, 0)
    if depth == math.MinInt {
        return 0
    }
    return depth
}
```

#### Описание алгоритма

Будем просто проходиться вниз и поддерживать максивальный depth смотря на уровень

#### О большое

Время: `O(N)`
Память: `O(H)`

### 100. Same Tree

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
        return p == nil && q == nil
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

#### Описание алгоритма

Первый случай очевидно на проверку значений дерева в одной точке если они не равны будем возвращать `false` далее смотрим если один из деревье равно `nil` то они оба должны быть `nil` а далее идем рекурсивно и проверяем чтобы и левые поддеревья были равны и правые.

#### О большое

Время: `O(N)`
Память: `O(H)`

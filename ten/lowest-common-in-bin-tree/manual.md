### 235. Lowest Common Ancestor of a Binary Search Tree

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root
    for node != nil {
        if p.Val > node.Val && q.Val > node.Val {
            node = node.Right
        } else if p.Val < node.Val && q.Val < node.Val {
            node = node.Left
        } else {
            return node
        }
    }
    return nil
}
```

### Описание алгоритма

Будем просто идти вниз до того момента пока не найдем родителя который будет

### O большое

Время: `O(N)`
Память: `O(H)`

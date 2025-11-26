### 572. Subtree of Another Tree

```go
func sameTree(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil || root2 == nil {
        return root1 == nil && root2 == nil
    }
    if root1.Val != root2.Val {
        return false
    }
    left := sameTree(root1.Left, root2.Left)
    right := sameTree(root1.Right, root2.Right)
    return left && right
}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    if sameTree(root, subRoot){
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
```
#### Описание алгоритма
Используем сравнение деревьев и проходимся вниз до того момента пока наш root не будет равен subRoot если дошли до конца вернем false

#### О большое
Время: `O(n)`
Память: `O(H)`
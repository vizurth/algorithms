### 124. Binary Tree Maximum Path Sum
```go
func dfs(root *TreeNode, sum *int) int {
    if root == nil {
        return 0
    }
    sumLeft := max(dfs(root.Left, sum), 0)
    sumRight := max(dfs(root.Right, sum), 0)
    *sum = max(*sum, root.Val + sumLeft + sumRight)
    return root.Val + max(sumLeft, sumRight)
}


func maxPathSum(root *TreeNode) int {
    maxSum := -1<<63
    dfs(root, &maxSum)
    return maxSum
}
```
### Описание алгоритма

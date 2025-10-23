### 144. Binary Tree Preorder Traversal
```go
func printTree(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    *result = append(*result, root.Val)
    printTree(root.Left, result)
    printTree(root.Right, result)
}

func preorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)

    printTree(root, &result)

    return result
}
```

### Описание алгоритма
Ну просто preorder чо такого ))))
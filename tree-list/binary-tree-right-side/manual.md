### 199. Binary Tree Right Side View
```go
func combineTree(root *TreeNode, result *[]int, level int){
    if root == nil {
        return 
    }
    if len(*result) < level + 1 {
        *result = append(*result, 0)
    }
    (*result)[level] = root.Val
    combineTree(root.Left, result, level + 1)
    combineTree(root.Right, result, level + 1)
}


func rightSideView(root *TreeNode) []int {
    result := make([]int, 0)
    combineTree(root, &result, 0)
    return result
}
```
### Описание алгоритма
ну проходимся по дереву в preorder и по кайфу
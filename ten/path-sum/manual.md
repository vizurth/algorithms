### 112. Path Sum

```go
func check(root *TreeNode, currSum int, targetSum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil && currSum + root.Val == targetSum {
        return true
    }
    left := check(root.Left, currSum + root.Val, targetSum)
    right := check(root.Right, currSum + root.Val, targetSum)
    return left || right
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    return check(root, 0, targetSum)
}
```

#### Описание алгоритма

Напишем функцию которая будет проходиться по дереву если дошли до пустого указателя возвращаем `nil`
Если дошли до листа дерева будем проверять не собрали ли мы нужную сумму, далее будем проверять левое и правое дерево на эту условие и если хоть в одном это будет работать то все ok

#### О большое

Время: `O(N)`
Память: `O(H)`

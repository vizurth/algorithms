### 21. Merge Two Sorted Lists
```go
func nodeToInt(node *ListNode) int {
    if node == nil {
        return math.MaxInt
    }
    return node.Val
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}

    curr := dummy

    for list1 != nil && list2 != nil {
        if nodeToInt(list1) <= nodeToInt(list2) {
            curr.Next = list1
            list1 = list1.Next
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next
    }

    for list1 != nil {
        curr.Next = list1
        list1 = list1.Next
    }

    for list2 != nil {
        curr.Next = list2
        list2 = list2.Next
    }

    return dummy.Next
}
```

#### Описание алгоритма
Просто берем и мирджим впринципе ничего нового
#### О большое 
Время: `O(N)`
Память: `O(1)`
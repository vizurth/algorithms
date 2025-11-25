### 83. Remove Duplicates from Sorted List

```go
func deleteDuplicates(head *ListNode) *ListNode {
    curr := head
    for curr != nil && curr.Next != nil {
        if curr.Val == curr.Next.Val {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return head
}
```

#### Описание алгоритма

Тут все просто просто проходимся по листу и сравниваем `curr.Val` c `curr.Next.Val` если они равны то просто убиваем связь `curr.Next = curr.Next.Next` если нет то просто идем дальше

#### О большое

Время: `O(N)`
Память: `O(1)`

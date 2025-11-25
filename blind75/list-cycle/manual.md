### 10. Linked List Cycle
```go
func hasCycle(head *ListNode) bool {
    fast := head
    slow := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            return true
        }
    }
    return false
}
```
### Описание алгоритма
Используем fast и slow если лист цикличный то fast может перегнать slow и когда нибудь они станут равны
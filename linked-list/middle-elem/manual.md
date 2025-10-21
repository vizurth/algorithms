### 876. Middle of the Linked List

```go

func middleNode(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

```

#### Описание алгоритма
Будем использовать Slow-Fast Pointer
slow += 1
fast += 2
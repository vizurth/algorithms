### 143. Reorder List

```go
func middlePlusElem(head *ListNode) *ListNode {
    fast := head
    slow := head
    for fast != nil && fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    return slow
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil

    curr := head
    for curr != nil {
        tmp := curr
        curr = curr.Next
        tmp.Next = prev
        prev = tmp
    }

    return prev
}

func reorderList(head *ListNode)  {
    middleElem := middlePlusElem(head)

    secondList := middleElem.Next
    middleElem.Next = nil
    secondList = reverseList(secondList)

    p1 := head
    p2 := secondList
    for p1 != nil && p2 != nil {
        p1_next := p1.Next
        p1.Next = p2
        p1 = p2
        p2 = p1_next
    }
}
```

#### Описание алгоритмов
возьмем разделим лист на два и разделим их и будем соединять по очереди
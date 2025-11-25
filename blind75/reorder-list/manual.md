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
        temp := curr
        curr = curr.Next
        temp.Next = prev
        prev = temp
    }

    return prev
}

func reorderList(head *ListNode) {
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
### Описание алгоритма
Для начала разделим наш лист на два отдельных. найдем middleElem - 1 и развернем вторую часть
далее будем перебрасывать указатели элементов которая написана в условии

### 206. Reverse Linked List

```go

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil

    curr = head

    for curr != nil{
        tmp = curr
        curr = curr.Next
        tmp.next = prev
        prev = temp
    }

    return prev
}

```

#### Описание алгоритма:

поставим `prev` указатель = `nil` чтобы привязать к нему первый элемент и оборвать связь с `Next` и будем двигать `curr` и привязывать его туда

#### Сложность алгоритма

`O(n)` - n кол-во элемента списка

#### Память

`O(1)` - так как все делается в одном массиве

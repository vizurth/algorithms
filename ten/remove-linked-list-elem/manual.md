### 203. Remove Linked List Elements
```go
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    curr := dummy

    for curr != nil && curr.Next != nil {
        if curr.Next.Val == val{
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }

    return dummy.Next
}
```
#### Описание алгоритма
Создаем `dummy` ноду чтобы не обрабатывать случай когда `head == nil` или когда все элементы равны `val` далее просто смотрим при проходе по листу на `curr.Next` и сравниваем его с `val` если это так то обрываем связь и прикрепляем к `curr.Next.Next` если нет то просто идем дальше
#### О большое
Время: `O(N)`
Память: `O(1)`
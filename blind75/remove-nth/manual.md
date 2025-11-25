### 19. Remove Nth Node From End of List
```go

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
    dummy := &ListNode{}
    dummy.Next = head
    curr := dummy
    
    for curr != nil{ // считаем длину
        length++
        curr = curr.Next
    }
    
    pointer := length - n - 1 // индекс указателя перед n-th с конца 
    curr = dummy
    for i := 0; i<pointer; i++{ // цикл дойдем до указателя индекс которого нашли
        curr = curr.Next
    }
    curr.Next = curr.Next.Next // обрываем связь с n-th
    
    return dummy.Next
}

```

#### Описание алгоритма
nil -> 1 -> 2 -> 3 -> 4 -> 5 -> nil
**dummy** - если будем удалять первый элемент из начала чтобы не вылезла ошибка
будем просто ставить `3.next = 3.next.next`, будем скипать => нужно поставить указатель на n + 1 c конца

#### Сложность алгоритма
`O(n)`, n - количество элементов списка

#### Оценка по памяти 
`O(1)` так как вы делаем все внутри уже существующего списка

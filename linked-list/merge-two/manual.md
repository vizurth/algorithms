### 21. Merge Two Sorted Lists
```go
func getValue(node *ListNode) int{
    if node == nil {
        return math.MaxInt
    }

    return node.Val
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}

    curr := dummy

    for list1 != nil || list2 != nil {
        if getValue(list1) > getValue(list2){
            curr.Next = list2
            list2 = list2.Next
        } else {
            curr.Next = list1
            list1 = list1.Next  
        }
    }

    return dumm
```
### Описание алгоритма
Алгоритм использует указатель curr для построения нового списка шаг за шагом, выбирая на каждом шаге меньший элемент из двух текущих узлов списков list1 и list2.

Чтобы упростить работу со списком, используется вспомогательный (фиктивный) узел dummy, от которого начинается результат.
В конце алгоритма возвращается dummy.Next, то есть начало объединённого списка.
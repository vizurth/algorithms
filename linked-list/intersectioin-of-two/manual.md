### 160. Intersection of Two Linked Lists
```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    dummy1 := headA
    dummy2 := headB
    for dummy1 != dummy2 {
        if dummy1 != nil {
            dummy1 = dummy1.Next
        } else {
            dummy1 = headA
        }
        if dummy2 != nil {
            dummy2 = dummy2.Next
        } else {
            dummy2 = headB
        }
    }

    return dummy1
}
```

### Описание алгоритма
Алгоритм использует два указателя, которые проходят по обоим спискам.
Ключевая идея:

Когда каждый указатель дойдет до конца своего списка, он начинает проходить другой список.

Таким образом, оба указателя проходят одинаковое общее расстояние (длину A + длину B), и если пересечение существует, они встретятся в одной и той же точке.
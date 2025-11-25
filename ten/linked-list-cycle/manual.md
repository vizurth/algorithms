### 141. Linked List Cycle
```go
func hasCycle(head *ListNode) bool {
    fast := head
    slow := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow{
            return true
        }
    }

    return false
}
```
### Описание алгоритма
Алгоритм использует классический метод “двух указателей” (алгоритм Флойда, Floyd’s Tortoise and Hare algorithm).

Идея проста:

Один указатель (slow) движется медленно — по одному шагу,
а второй (fast) — быстро — по два шага.
Если в списке есть цикл, то в какой-то момент они встретятся внутри цикла.

### О большое:
Время: `O(N)` - даже без цикла мы будем идти указателями не больше N
Память: `O(1)`
### 234. Palindrome Linked List

```go

func middleElem(head *ListNode) *ListNode {
    slow := head
    fast := head
    
    for fast != nil && fast.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    curr := head
    
    for curr != nil{
        tmp := curr
        curr = curr.Next
        tmp.Next = prev
        prev = tmp
    }
    
    return prev
}
 
func isPalindrome(head *ListNode) bool {
    middle := middleElem(head)
    reverseList := reverseList(middle)
    
    p1 := head
    p2 := reverseList
    
    for p1 != nil && p2 != nil {
        if p1.Val != p2.Val{
            return false
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    
    return true
}

```

#### Описание алгоритмов
Найдем *middleElem* и развернем оставшуюся часть листа после него.
Поставим d и p2, и будем сравнивать Val до момента пока один из них неравен nil
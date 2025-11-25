### 23. Merge k Sorted Lists
```go
func findMinIdx(lists []*ListNode) (int, bool) {
    minIdx, minVal := math.MaxInt, math.MaxInt

    for idx, list := range lists {
        if list == nil {
            continue
        }
        if list.Val < minVal {
            minIdx = idx
            minVal = list.Val
        }
    }


    if minVal == math.MaxInt {
        return 0, false
    }

    return minIdx, true
}

func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{}

    curr := dummy

    for {
        idx, ok := findMinIdx(lists)
        if !ok {
            curr.Next = nil 
            break
        }
        curr.Next = lists[idx]
        lists[idx] = lists[idx].Next
        
        curr = curr.Next
    }

    return dummy.Next
}
```
### Описание алгоритмов

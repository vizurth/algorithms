### 23. Merge k Sorted Lists
```go
func getMinIdx(lists []*ListNode) (int, bool){
    minVal, minIdx := math.MaxInt, math.MaxInt
    
    for idx, node := range lists {
        if node == nil{
            continue
        }
        if node.Val < minVal {
            minIdx = idx
            minVal = node.Val
        }
    }
    
    if minIdx == math.MaxInt{
        return 0, false
    }
    
    return minIdx, true
}


func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{}
    
    curr := dummy
    
    for {
        minIdx, ok := getMinIdx(lists)
        if !ok {
            curr.Next = nil
            break
        }
        curr.Next = lists[minIdx]
        lists[minIdx] = lists[minIdx].Next
        
        curr = curr.Next
    }
    
    return dummy.Next
}
```

### Описание алгоритма
Алгоритм выполняет итеративное построение нового списка, на каждом шаге выбирая минимальный элемент среди текущих голов всех списков.

Для поиска минимального узла используется вспомогательная функция getMinIdx(), которая находит индекс списка с наименьшим текущим элементом.

Таким образом, алгоритм последовательно «вытягивает» минимальные значения из всех списков, пока все они не будут полностью обработаны.
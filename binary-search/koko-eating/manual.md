### 875. Koko Eating Bananas
```go
func check(k int, piles []int, h int) bool {
	currHours := 0
	for _, count := range piles {
		currHours += (count + k - 1) / k
	}
	return currHours <= h
}


func minEatingSpeed(piles []int, h int) int {
    l, r := 1, 0
    for _, pile := range piles {
        if pile > r {
            r = pile
        }
    }
    
    for l < r {
        x := (l + r) / 2
        if check(x, piles, h) {
            r = x
        } else {
            l = x + 1
        }
    }
    return l
}
```
### Описание алгоритма

### 121. Best Time to Buy and Sell Stock
```go
func maxProfit(prices []int) int {
    i, j := 0, 1
    res := 0
    for j < len(prices) {
        curr := prices[j] - prices[i]
        if prices[i] < prices[j] {
            res = max(res, curr)
        } else {
            i = j
        }
        j++
    }
    return res
}
```
### Описание алгоритма
Мое решение основано на использовании двух указателесь вы двигаем правый до самого конца, а левый если он будет больше чем правый, и поддерживаем максимальное число

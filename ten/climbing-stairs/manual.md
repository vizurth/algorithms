### 70. Climbing Stairs
```go
func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[0] = 0
    dp[1] = 1
    if n >= 2 {
        dp[2] = 2
    }
    for i := 3; i < len(dp); i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```
### Описание алгоритма:
Тут используем просто ДП с одной переменнтой и прибавляем к текущей позиции, позиции на две и на одну ступеньки меньше.
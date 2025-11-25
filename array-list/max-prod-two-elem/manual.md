### 1464. Maximum Product of Two Elements in an Array
```go
func maxProduct(nums []int) int {
    max1, max2 := math.MinInt, math.MinInt
    for _, elem := range nums {
        if elem > max1 {
            max2 = max1
            max1 = elem
        } else if elem <= max1 && elem > max2 {
            max2 = elem
        }
    }

   return (max1 - 1) * (max2 - 1) 
}
```
### Описание алгоритма
Будем держать два максимума на массиве и потом умножать их и все

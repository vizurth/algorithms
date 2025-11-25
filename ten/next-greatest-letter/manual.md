### 744. Find Smallest Letter Greater Than Target

```go
func nextGreatestLetter(letters []byte, target byte) byte {
    left, right := 0, len(letters)-1

    if target >= letters[right] {
        return letters[0]
    }

    for left < right {
        m := left + (right-left)/2
        if letters[m] > target {
            right = m
        } else {
            left = m + 1
        }
    }

    return letters[left]
}

```

#### Описание алгоритма

Используем бинпоиск и проверяем что target >= последнего элемента тогда просто пикаем первые элемент массива

#### О большое

Время: `O(log N)`
Память: `O(1)`

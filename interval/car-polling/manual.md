### 1094. Car Pooling
```go
func carPooling(trips [][]int, capacity int) bool {
    points := make([][]int, 0)
    for _, elem := range trips {
        points = append(points, []int{elem[1], elem[0]})
        points = append(points, []int{elem[2], -1 * elem[0]})
    }
    // очень важно сортировку проверять
    sort.Slice(points, func(a, b int) bool {
    if points[a][0] == points[b][0] {
        return points[a][1] < points[b][1] // сначала высадка (-), потом посадка (+)
    }
    return points[a][0] < points[b][0]
})
    countSum := 0
    for _, elem := range points {
        countSum += elem[1]
        if countSum > capacity {
            return false
        }
    }
    return true
}
```
### Описание алгоритма
Используем подход "линейного сканирования" (line sweep) — такой же, как при подсчёте перекрывающихся интервалов:

каждое событие (вход и выход пассажиров) добавляем как "точку на линии";

сортируем все события по координате;

идём слева направо, прибавляя или вычитая количество пассажиров;

если в любой момент число пассажиров превышает capacity — возвращаем false.

просто умножаем -1 * кол-во машин и +1 * кол-во машин
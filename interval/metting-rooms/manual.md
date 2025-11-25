### Metting Rooms
```go

func solve1(A [][]int) int { // Время: O(N*logN) Память: O(N)
    points := make([][]int, 0, len(A)*2)

    // Формируем "точки событий"
    for _, interval := range A {
        points = append(points, []int{interval[0], +1}) // заселение
        points = append(points, []int{interval[1], -1}) // выселение
    }

    // Сортируем события
    // Если время одинаковое, сначала идет -1 (выселение), потом +1 (заселение)
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] {
            return points[i][1] < points[j][1]
        }
        return points[i][0] < points[j][0]
    })

    maxRooms := 0
    currRooms := 0

    // Подсчёт количества используемых комнат во времени
    for _, p := range points {
        currRooms += p[1]
        if currRooms > maxRooms {
            maxRooms = currRooms
        }
    }

    return maxRooms
}

// Если надо найти было ли пересечение O(N) O(1)
func solve2(A [][]int) bool {
	sort.Slice(A, func(i, j int) {
		return A[i][0] < A[j][0]
	})
	for i:=0; i<len(A)-1; i++ {
		if A[i][1] < A[i+1][1]{
			return false
		}
	}
	return false
}

```
### Описание алгоритма
Будем использовать метод точек
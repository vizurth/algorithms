### 1024. Video Stiching
```go
func videoStitching(clips [][]int, time int) int {
    sort.Slice(clips, func (a, b int) bool {
        return clips[a][0] < clips[b][0]
    })

    n := len(clips)
    currEnd, farCanReach, total := 0,0,0
    i := 0

    for currEnd < time {
        total++
        for i < n && currEnd >= clips[i][0] {
            farCanReach = max(farCanReach, clips[i][1])
            i++
        }
        if farCanReach == currEnd {
            return -1
        }
        currEnd = farCanReach
    }

    return total
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
### Описание алгоритма
Это жадный алгоритм (greedy):
На каждом шаге выбираем максимально дальний конец видео, который можно достичь из текущего диапазона.

Аналогия:
Представь, что ты идёшь по дороге длиной time.
Каждый клип — это участок, который можно пройти.
Ты хочешь дойти до конца дороги, используя минимальное количество шагов, при этом каждый шаг выбираешь так, чтобы продвинуться как можно дальше.
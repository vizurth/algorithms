### 4. Longest Palindromic Substring
```go
func longestPalindrome(s string) string {
    if len(s) < 2 {return s}
    res := fmt.Sprintf("%c", s[0])

    m := make(map[byte][]int)
    
    for i := 0; i < len(s); i++ {
        m[s[i]] = append(m[s[i]], i)
        
        if len(m[s[i]]) > 1 {
            for _, val := range m[s[i]] {
                if i != val && isPalindrome(s[val:i+1]) {
                    if i - val + 1 > len(res) {
                        res = s[val:i+1]
                    }
                }
            }
        }
    }

    return res
}

func isPalindrome(s string) bool {
    if len(s) == 1 {
        return true
    }
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }

    return true
}
```
### Описание алгоритма
Для начала напишем функцию которая будет определять является ли слово палиндромом для этого будем идти от l = 0, r = len(s) - 1 и будет сравнивать `s[l] и s[r]` если они не равны возвращаем false

Далее перейдем к нашему алгоритму. Будем хранить индексы букв которые нам встречаются в map. далее если бук в списке появляется больше чем 1 можно начинать и проверять слово по началу индекса и концу на полиндром
и уже если длина оказывает больше будем обновлять результат
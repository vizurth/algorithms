### 17. Valid Parentheses
```go
type Stack []string

func (s *Stack) Pop() string {
    if len(*s) == 0 {
        return "false"
    }
    res := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return res
}

func isValid(s string) bool {
    if len(s) < 2 {
        return false
    }

    stack := make(Stack, 0)

    for i := 0; i < len(s); i++ {
        elem := string(s[i])
        if elem == "(" || elem == "{" || elem == "[" {
            stack = append(stack, elem)
        } else if elem == ")" {
            prev := stack.Pop()
            if prev != "(" || prev == "false" {
                return false
            }
        } else if elem == "}" {
            prev := stack.Pop()
            if prev != "{" || prev == "false" {
                return false
            }
        } else if elem == "]" {
            prev := stack.Pop()
            if prev != "[" || prev == "false" {
                return false
            }
        }
    }
    if len(stack) == 0 {
        return true
    } else {
        return false
    }
}
```
### Описание алгоритма
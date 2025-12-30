package validparentheses

func isValid(s string) bool {
    bracketMap := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    stack := make([]rune, 0, len(s))
    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
        } else {
            open := bracketMap[ch]
            if len(stack) == 0 || stack[len(stack) - 1] != open {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

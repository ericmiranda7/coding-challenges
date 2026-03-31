func isValid(s string) bool {
    stk := []rune{}
    cmap := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, r := range s {
        if _, ok := cmap[r]; !ok {
            stk = append(stk, r)
        } else {
            if len(stk) == 0 {
                return false
            }
            popped := stk[len(stk)-1]
            stk = stk[:len(stk)-1]
            if popped != cmap[r] {
                return false
            }
        }
    }

    return len(stk) == 0 && true
}

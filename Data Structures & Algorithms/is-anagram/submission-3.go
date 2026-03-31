func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := map[rune]int{}

    for _, c := range s {
        count[c] += 1
    }

    for _, c := range t {
        count[c] -= 1
    }

    for _, v := range count {
        if v != 0 {
            return false
        }
    }

    return true
}

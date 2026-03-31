func letterCombinations(digits string) []string {
    res := []string{}
    digiMap := map[rune]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    combo(digits, digiMap, 0, "", &res)
    return res
}

func combo(digits string, digiMap map[rune]string, i int, substring string, res *[]string) {
    if i == len(digits) {
        if len(substring) == len(digits) && substring != "" {
            *res = append(*res, substring)
        }
        return
    }

    runes := []rune(digits)
    chars := digiMap[runes[i]]

    for _, c := range []rune(chars) {
        substring += string(c)
        combo(digits, digiMap, i+1, substring, res)
        substring = substring[:len(substring)-1]
    }
}

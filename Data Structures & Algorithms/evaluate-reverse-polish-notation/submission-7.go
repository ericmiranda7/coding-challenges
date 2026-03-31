
func evalRPN(tokens []string) int {
    s := []string{}
    optMap := map[string]bool{
        "+": true,
        "-": true,
        "*": true,
        "/": true,
    }

    for _, op := range tokens {
        if optMap[op] {
            // operator
            ops := s[len(s) - 2:]
            op1, _ := strconv.Atoi(ops[0])
            op2, _ := strconv.Atoi(ops[1])
            operands := []int{op1, op2}
            s = s[:len(s)-2]

            var res int
            switch op {
            case "+":
                res = operands[0] + operands[1]
            case "-":
                res = operands[0] - operands[1]
            case "*":
                res = operands[0] * operands[1]
            case "/":
                res = operands[0] / operands[1]
            default:
                fmt.Println("huh")
            }

            resStr := strconv.Itoa(res)
            s = append(s, resStr)
        } else {
            // operand
            s = append(s, op)
        }
    }

    res, _ := strconv.Atoi(s[0])

    return res
}

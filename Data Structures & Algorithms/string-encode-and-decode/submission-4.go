type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    sb := strings.Builder{}

    for _, s := range strs {
        slen := len(s)

        _, err := sb.WriteString(fmt.Sprintf("%v!%v", slen, s))
        if err != nil {
            fmt.Println(err)
        }
    }

    fmt.Println("no errs")

    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    var res []string

    for len(encoded) > 0 {
        bangIndex := strings.Index(encoded, "!")
        slen, _ := strconv.Atoi(encoded[:bangIndex])
        s := encoded[bangIndex+1:bangIndex+1+slen]
        res = append(res, s)
        encoded = encoded[bangIndex+1+slen:]
    }
    fmt.Println("reached")

    return res
}

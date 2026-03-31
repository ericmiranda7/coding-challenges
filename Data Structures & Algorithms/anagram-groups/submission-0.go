func groupAnagrams(strs []string) [][]string {
    amap := map[[26]int][]string{}

    for _, s := range strs {
        arep := [26]int{}
        for _, r := range s {
            arep[int(r)-'a'] += 1
        }

        if slist, ok := amap[arep]; ok {
            amap[arep] = append(slist, s)
        } else {
            amap[arep] = []string{s}
        }
    }

    res := [][]string{}
    for _, v := range amap {
        res = append(res, v)
    }

    return res
}

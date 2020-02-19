func compare() {
    if len(os.Args[1:]) != 2 {
        z01.PrintRune('\n')
        return
    }
    a := os.Args[1]
    b := os.Args[2]
    c := 0
    if a < b {
        c = -1
    } else if a > b {
        c = 1
    }
    res := strconv.Itoa(c)
    for _, v := range res {
        z01.PrintRune(v)
    }
}
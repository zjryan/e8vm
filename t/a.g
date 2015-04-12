func fabo(n int) int {
    if n == 0 return 0
    if n == 1 return 1
    return fabo(n-1) + fabo(n-2)
}

func main() {
    var a int
    for a != 20 {
        printInt(fabo(a))
        a = a + 1
    }
}

func fabo(n int) int {
    if n == 0 return 0
    if n == 1 return 1
    return fabo(n-1) + fabo(n-2)
}

func main() {
    var i int
    for i < 3 {
        var a int
        for a != 20 {
            printInt(fabo(a))
            a = a + 1
            if a <= 5 { continue }
            break
        }
        printInt(1000)
        i = i + 1
    }
}

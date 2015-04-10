func p(a int, b int) {
	printInt(a)
	printInt(b)
}

func main() {
    var a int
    for a != 3 {
        p(a, a)
        a = a + 1
    }
}

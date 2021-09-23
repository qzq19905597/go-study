package main
import "fmt"
const LIMIT = "limit"
const (
	a = 1
	b = 2
	c = 3
)
const (
	d = iota
	e
	f
	h = "h"
	g = iota
)

func main() {
	fmt.Println(LIMIT)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f, h, g)
}
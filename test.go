package main
import "fmt"
func main (){
	// var a int
	// var b string
	// var c bool
	// var d = [2]{1, 2}
	// var e = d
	_, f, h := numbers()
	// fmt.Println(a, b, c, d, e)
	// fmt.Println("hello world")
	fmt.Println(f, h)
}

func numbers()(int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
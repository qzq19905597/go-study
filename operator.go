//基本运算符与js相同

// （& *）运算符使用
package main

import "fmt"

func main() {
	var a = 3;
	var b = func() {

	}
	var ptr = &a;
	var ptr2 = &b;
	
	fmt.Println("a:", *ptr);
	fmt.Println("b:", *ptr2);
}
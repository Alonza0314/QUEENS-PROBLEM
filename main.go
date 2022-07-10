package main

/*
#include "c_function.c"
void clear_screen();
void sys_pause();
*/
import "C"
import (
	"fmt"

	"QUEEN/function"
)

var n int
var arr [][]int
var num int = 0
var pr int

func main() {
	C.clear_screen()
	fmt.Println("Input the size of the board:(1~10)")
	fmt.Scan(&n)
	fmt.Println("The size of the board is", n)

	for i := 0; i < n; i++ {
		in_temp := make([]int, n)
		arr = append(arr, in_temp)
	}

	fmt.Println("Do yu want to print the playboard?")
	fmt.Print("1:yes\n2:no\n")
	fmt.Scan(&pr)

	output := function.Solve(n, 0, arr, num, pr)
	fmt.Println("total num =", output)

	C.sys_pause()
}

package function

import "fmt"

func print(n int, num int, arr [][]int) {
	fmt.Println("num =", num)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println(arr[i][n-1])
	}
	fmt.Println()
}

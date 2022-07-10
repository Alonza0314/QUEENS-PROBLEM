package function

func Solve(n int, row int, arr [][]int, num int, pr int) (num_ int) {
	for i := 0; i < n; i++ {
		arr[row][i] = 1
		if check(n, row, i, arr) == 1 {
			if row == n-1 {
				num++
				if pr == 1 {
					print(n, num, arr)
				}
			} else {
				num = Solve(n, row+1, arr, num, pr)
			}
		}
		arr[row][i] = 0
	}
	num_ = num
	return
}

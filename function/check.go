package function

func check(n int, row int, col int, arr [][]int) (ret int) {
	if row == 0 {
		ret = 1
		return
	}
	for i := 0; i < row; i++ {
		if arr[i][col] == 1 {
			ret = 0
			return
		}
	}

	var i int = row - 1
	var j int = col - 1
	for {
		if i < 0 || j < 0 {
			break
		}
		if arr[i][j] == 1 {
			ret = 0
			return
		}
		i--
		j--
	}
	i = row - 1
	j = col + 1
	for {
		if i < 0 || j >= n {
			break
		}
		if arr[i][j] == 1 {
			ret = 0
			return
		}
		i--
		j++
	}
	ret = 1
	return
}

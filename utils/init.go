package utils

func Init2DArray(row int, col int) (arr [][]int) {
	arr = make([][]int, row)

	for i := 0; i < col; i++ { 
		arr[i] = make([]int, col) 
	}

	return
}
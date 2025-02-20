package lat2

import (
	"fmt"

	"github.com/mnqrt/latihan-logic/utils"
)

func No1(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 1)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			arr[row][col] = 2 * col + 1
		}
	}

	return
}

func No2(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 2)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			arr[row][col] = 2 * col + 2
		}
	}

	return
}

func No3(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 3)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			arr[row][col] = 2 * n * row + 2 * col + 1
		}
	}

	return
}

func No4(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 5)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			if row == 0 {
				arr[row][col] = 1 + col * 3
			} else if col <= 1 {
				arr[row][col] = arr[row - 1][n - 1] + 3 * (col + 1)
			} else {
				arr[row][col] = arr[row][col - 1] + 2
			}
		}
	}

	return
}

func No6(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 6)
	arr = utils.Init2DArray(n, n)
	cur := 1

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row % 2 == 0 {
				arr[row][col] = cur
				cur += 3
			} else {
				arr[row][n - 1 - col] = cur
				cur += 2
			}
		}
	}
	return arr
}

func No9(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 9)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			if row == col || row + col + 1 == n { 
				arr[row][col] = 2 * col + 1 
			} 
		}
	}

	return
}


func No10(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 10)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			if row >= col { 
				arr[row][col] = 2 * col + 1 
			} 
		}
	}

	return
}

func No12(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 12)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			if (col <= row && row + col + 1 <= n) ||
					(col >= row && row + col + 1 >= n) { 
				arr[row][col] = 2 * col + 1 
			} 
		}
	}

	return
}

func No13(n int) (arr [][]int) {
	fmt.Printf("This is no #%d\n", 13)
	arr = utils.Init2DArray(n, n)

	for row := 0; row < n; row ++ {
		for col := 0; col < n; col ++ {
			if (col >= row && row + col + 1 <= n) ||
					(col <= row && row + col + 1 >= n) { 
				arr[row][col] = 2 * col + 1 
			} 
		}
	}

	return
}
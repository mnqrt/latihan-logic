package lat3

import (

	"github.com/mnqrt/latihan-logic/utils"
)

func No1(n int) (result [][]int) {
	result = utils.Init2DArray(n, n)
	value := 1

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i % 2 == 0 {
				result[i][j] = value
			} else { result[i][i - j] = value }
			value += 2
		}
	}
	return result
}

func No2(n int) (result [][]int) {
	result = utils.Init2DArray(n, n)
	value := 1

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i % 2 == 0 {
				result[i][j] = value
			} else { result[i][n - (j - i) - 1] = value }
			value += 2
		}
	}
	return result
}

func No3(n int) [][]int {
	result := utils.Init2DArray(n, n)
	value := -2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > n - 1 - i {
				continue
			}

			if i == 0 {
				value += 3
				result[i][j] = value
				continue
			}

			if i % 2 == 0 {
				value += 2
				result[i][j] = value
			} else {
				value += 3
				result[i][n - i - j - 1] = value
			}
		}  
	}
	return result
}



func No4(n int) [][]int {
	result := utils.Init2DArray(n, n)
	value := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > i {
				continue
			}

			if i % 2 == 0 {
				result[i][n - j - 1] = value
			} else {
				result[i][n - i + j - 1] = value
			}

			value += 2
		}
	}

	return result
}

func No5(n int) [][]int {
	result := utils.Init2DArray(n, n)
	val := 1
	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		for j := 0; j <= mid; j++ {
			if (i < j || i+j > n-1) {
				continue
			}

			if i%2 == 0 {
				result[i][j] = val
				result[i][n - j - 1] = val
				result[n - i - 1][n - j - 1] = val
				result[n - i - 1][j] = val
			} else {
				result[i][i - j] = val
				result[i][n - i + j - 1] = val
				result[n - i - 1][n - i + j - 1] = val
				result[n - i - 1][i - j] = val
			}

			val += 2
		}
	}

	return result
}

func No6(n int) [][]int {
	result := utils.Init2DArray(n, n)
	val := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > j || i + j > n - 1 {
				continue
			}

			if i%2 == 0 {
				result[i][j] = val
				result[n - i - 1][j] = val
			} else {
				result[i][n - j - 1] = val
				result[n - i - 1][n - j - 1] = val
			}

			val += 2
		}
	}

	return result
}

func No7(n int) [][]int {
	result := utils.Init2DArray(n, n)
	val := 1

	for i, row := range result {
		if i > (n - 1) / 2 {
			break
		}

		for j := range row {
			if j + i < (n - 1) / 2 || j - i > (n - 1) / 2 {
				continue
			}

			if i % 2 != 0 {
				result[i][j] = val
				result[n - i - 1][j] = val
			} else {
				result[i][n - j - 1] = val
				result[n - i - 1][n - j - 1] = val

			}

			val += 2
		}
	}

	return result
}

func No8(n int) [][]int {
	result := utils.Init2DArray(n, n)
	val := 1
	mid := (n - 1) / 2

	for i := 0; i < n; i++ {
		if i > mid {
			break
		}
		for j := 0; j < n; j++ {
			if j + i < mid || j - i > mid {
				continue
			}

			if i%2 != 0 {
				result[j][i] = val
				result[j][n - i - 1] = val
			} else {
				result[n - j - 1][i] = val
				result[n - j - 1][n - i - 1] = val
			}

			val += 2
		}
	}

	return result
}

func No9(n int) [][]int {
	result := utils.Init2DArray(n, n)
	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		val := 1
		for j := 0; j <= mid; j++ {
			if i >= mid - j && i + j >= mid {
				result[i][j] = val
				result[i][n - 1 - j] = val
				result[n - 1 - i][j] = val
				result[n - 1 - i][n - 1 - j] = val

				val += 2
			}

		}
	}

	return result
}


func No10(n int) [][]int {
	result := No9(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = n + 1 - result[i][j]
		}
	}

	return result
}


func No11(n int) [][]int {
	result := utils.Init2DArray(n, n)
	val := 1
	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		for j := 0; j <= mid; j++ {
			if j > i || i + j > n - 1 {
				continue
			}
			
			if i%2 == 0 {
				result[i][n - 1 - j] = val
				result[n - 1 - i][n - 1 - j] = val
			} else {
				result[i][n - 1 + j - i] = val
				result[n - 1 - i][n - 1 + j - i] = val
			}
			
			if j == i && i < mid {
				result[i][j] = i * 2 + 1
				result[n - 1 - i][j] = i * 2 + 1
			}

			val += 2
		}
	}

	return result
}

func No11B(n int) [][]int {
	mid := (n - 1) / 2
	result := utils.Init2DArray(n, n + mid)
	val := 1

	for i := 0; i <= mid; i++ {
		for j := 0; j <= mid; j++ {
			if j > i || i + j > n - 1 {
				continue
			}
			
			if i%2 == 0 {
				result[i][n - 1 - j] = val
				result[i][n - 1 + j] = val

				result[n - 1 - i][n - 1 - j] = val
				result[n - 1 - i][n - 1 + j] = val
			} else {
				result[i][n - 1 + j - i] = val
				result[i][n - 1 - j + i] = val

				result[n - 1 - i][n - 1 +j - i] = val
				result[n - 1 - i][n - 1- j + i] = val
			}
			
			if j == i && i < mid {
				result[i][j] = i * 2 + 1
				result[n - 1 - i][j] = i * 2 + 1
			}

			val += 2
		}
	}

	return result
}

func No12(n int) [][]int {
	result := utils.Init2DArray(n, n)
	prev := No11(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = prev[i][n - j - 1]
		}
	}

	return result
}

func No12B(n int) [][]int {
	mid := (n - 1) / 2
	result := utils.Init2DArray(n, n + mid)
	prev := No11B(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n + mid; j++ {
			result[i][j] = prev[i][n + mid - j - 1]
		}
	}

	return result
}


func No14(n int) [][]int {
	cur := 1
	prev := 0
	result := utils.Init2DArray(n, n)

	for col := 0; col < n; col ++ {
		for row := 0; row < n; row ++ {
			if col % 2 == 0 {
				result[row][col] = cur
			} else {
				result[n - 1 - row][col] = cur
			}

			new := cur + prev
			prev = cur
			cur = new
		}
	}

	return result
}
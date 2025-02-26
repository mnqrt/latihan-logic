package lat1

import "fmt"

func No1(n int) (arr []int) {
	fmt.Printf("This is no #%d\n", 1)
	cur := 1
	prev := 0;
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = cur
		new := cur + prev
		prev = cur
		cur = new
	}
	return
}

func No5(n int) (arr []int) {
	fmt.Printf("This is no #%d\n", 5)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 2 * n - i * 2;
	}
	return
}

func No9(n int) (arr []int) {
	fmt.Printf("This is no #%d\n", 9)

	arr = make([]int, n)
	for i := 0; i < n; i++ {
		if i < n / 2 {
			arr[i] = 3 * i + 3
		} else if 2 * i + 1 == n {
			arr[i] = arr[i - 1] + 2
		} else {
			arr[i] = 3 * (n / 2) - 3 * (i - (n + 1) / 2)
		}
	}
	return
}
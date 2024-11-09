package main

import "fmt"

var arr = []int{1}
var n int

func triangle_Pa(arr []int) []int {
	var arr_ans = []int{}
	arr_ans = append(arr_ans, 1)
	for i := 0; i < len(arr)-1; i++ {
		arr_ans = append(arr_ans, arr[i]+arr[i+1])
	}
	arr_ans = append(arr_ans, 1)
	return arr_ans
}

func var_dump(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var_dump(arr)
		arr = triangle_Pa(arr)
	}
	var_dump(arr)
}

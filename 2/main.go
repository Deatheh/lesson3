package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var str string
var ans int

func chet(str string) string {
	var new_arr = []string{}
	arr := strings.Fields(str)
	fmt.Println(arr[1])
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] == "*" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			if 1 == 0 {
				print(err)
			}
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			new_arr = append(new_arr, fmt.Sprint(tmp_1*tmp_2))
		} else if arr[i] == "/" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			if 1 == 0 {
				print(err)
			}
			new_arr = append(new_arr, fmt.Sprint(tmp_1/tmp_2))
		} else {
			new_arr = append(new_arr, arr[i])
		}
	}
	/*arr = new_arr
	new_arr = []string{}
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] == "+" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			if 1 == 0 {
				print(err)
			}
			new_arr = append(new_arr, fmt.Sprint(tmp_1+tmp_2))
		} else if arr[i] == "-" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			if 1 == 0 {
				print(err)
			}
			new_arr = append(new_arr, fmt.Sprint(tmp_1-tmp_2))
		}
	}*/

	for i := 0; i < len(new_arr); i++ {
		fmt.Println(new_arr[i])
	}
	return new_arr[0]
}

func findScope(str string) string {
	var new_str, new_str_tmp string
	var key = true
	for key {
		key = false
		for i := 0; i < len(str); i++ {
			if new_str_tmp == "" {
				if str[i] != '(' {
					new_str += string(str[i])
				} else {
					key = true
					new_str_tmp = "("
				}
			} else {
				if str[i] == ')' {
					new_str += chet(new_str_tmp[1:])
					new_str_tmp = ""
				} else if str[i] == '(' {
					new_str += new_str_tmp
					new_str_tmp = "("
				} else {
					new_str_tmp += string(str[i])
				}
			}
		}
		str = new_str
		new_str = ""
		new_str_tmp = ""
	}
	return chet(new_str)
}

func main() {
	file, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	str = string(file)
	fmt.Println(str)
	//fmt.Println(chet(str))
}

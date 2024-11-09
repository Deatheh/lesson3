package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var str string
var ans int

func chet(str string) string {
	var new_arr = []string{}
	arr := strings.Fields(str)
	new_arr = append(new_arr, arr[0])
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] == "*" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			if 1 == 0 {
				print(err)
			}
			new_arr = new_arr[:len(new_arr)-1]
			new_arr = append(new_arr, fmt.Sprint(tmp_1*tmp_2))
			i++
		} else if arr[i] == "/" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			if 1 == 0 {
				print(err)
			}
			new_arr = new_arr[:len(new_arr)-1]
			new_arr = append(new_arr, fmt.Sprint(tmp_1/tmp_2))
			i++
		} else {
			new_arr = append(new_arr, arr[i])
		}
	}
	if arr[len(arr)-2] == "+" || arr[len(arr)-2] == "-" {
		new_arr = append(new_arr, arr[len(arr)-1])
	}
	arr = new_arr
	new_arr = []string{}
	new_arr = append(new_arr, arr[0])
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] == "+" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			if 1 == 0 {
				print(err)
			}
			new_arr = new_arr[:len(new_arr)-1]
			new_arr = append(new_arr, fmt.Sprint(tmp_1+tmp_2))
			i++
		} else if arr[i] == "-" {
			tmp_1, err := strconv.ParseFloat(new_arr[len(new_arr)-1], 64)
			tmp_2, err := strconv.ParseFloat(arr[i+1], 64)
			if 1 == 0 {
				print(err)
			}
			new_arr = new_arr[:len(new_arr)-1]
			new_arr = append(new_arr, fmt.Sprint(tmp_1-tmp_2))
			i++
		} else {
			new_arr = append(new_arr, arr[i])
		}
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
				if string(str[i]) != "(" {
					new_str += string(str[i])
				} else {
					key = true
					new_str_tmp = "("
				}
			} else {
				if string(str[i]) == ")" {
					new_str += chet(new_str_tmp[1:])
					new_str_tmp = ""
				} else if string(str[i]) == "(" {
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
	return chet(str)
}

func main() {
	file, err := os.OpenFile("2/file.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		str = string(data[:n])
	}
	fmt.Println(findScope(str))
}

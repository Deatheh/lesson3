package main

import (
	"fmt"
	"strings"
)

var prefix, tmp string
var n int

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if i == 0 {
			prefix = tmp
			continue
		}
		for j := 0; j <= len(prefix); j++ {
			if j == len(prefix) {
				prefix = ""
				break
			}
			if strings.Contains(tmp, prefix[:len(prefix)-j]) {
				prefix = prefix[:len(prefix)-j]
				break
			}
		}
	}
	if prefix == "" {
		fmt.Println("Prefix no find!!!")
	} else {
		fmt.Println(prefix)
	}

}

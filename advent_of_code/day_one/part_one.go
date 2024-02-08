package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileData, err := os.ReadFile("input.txt")
	check(err)

	data := string(fileData)
	data_arr := strings.Split(data, "\n")

	var res int
	numb_list := "1234567890"

	for i := 0; i < len(data_arr); i++ {
		left_numb := ""
		right_numb := ""

		for j := 0; j < len(data_arr[i]); j++ {
			found_numb := strings.Contains(numb_list, string(data_arr[i][j]))
			if found_numb {
				right_numb = string(data_arr[i][j])
				if len(left_numb) == 0 {
					left_numb = string(data_arr[i][j])
				}
			}
		}

		numb, _ := strconv.Atoi(left_numb + right_numb)
		res += numb
	}
	fmt.Println(res)
}

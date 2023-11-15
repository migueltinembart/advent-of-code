package main

import (
	"fmt"
	"strconv"
)

func main() {
	list := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}
	elvMap := mapElves(list)

	fmt.Println(elvMap)

}

func mapElves(list []string) map[int]int {
	m := make(map[int]int)

	iterator := 1
	for _, v := range list {
		num, err := strconv.Atoi(v)
		if err != nil {
			iterator++
		}
		m[iterator] += num
	}
	return m
}


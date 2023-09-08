package main

import "fmt"

func main() {

	var pMap = make(map[int]bool, 10)
	var s = []int{1, 2, 3, 4}
	for _, b := range s {
		pMap[b] = false
	}
	fmt.Println(pMap)
}

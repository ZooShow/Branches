package main

import "fmt"

func max(args ...int) int {
	m := args[1]
	// v:= range args
	for v, _ := range args {
		if args[v] > m {
			m = args[v]
		}
	}
	return m
}

func main() {
	masive := []int{15, 28, 3, 4, 5, 6}
	fmt.Println(max(masive...))
}

package main

import "fmt"

func main() {
	mas := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, v := range mas {
		set[v] = struct{}{}
	}

	res := make([]string, 0, len(set))
	for k, _ := range set {
		res = append(res, k)
	}
	fmt.Println(res)
}

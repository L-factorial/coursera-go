package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := make([]int, 0, 3)
	var item string
	for {
		fmt.Println("Enter the item ( Enter X to terminate")
		fmt.Scan(&item)
		if(strings.HasPrefix(item, "X")){
			fmt.Println("Exiting ...")
			break
		}
		v,_ := strconv.Atoi(item)
		s = append(s, v)
		sort.Ints(s)

		fmt.Printf("%v\n", s)
	}
}

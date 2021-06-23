package main

import (
	"fmt"
)
func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if s[j+1] < s[j] {
				Swap(s, j) // swap the item s[j-1] and s[j]
			}
		}
	}
}
//Swap the item in slice s between the indices idx-1 and idx
func Swap(s []int, idx int) {
	temp := s[idx+1]
	s[idx+1] = s[idx]
	s[idx] = temp
}
func main() {
	fmt.Println("Enter the total number of items in the array to be sorted : ")
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Exiting ...")
		return
	}
	//fmt.Printf("%d", n)
	fmt.Println()
	var num int
	s := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Printf("arr[%d]=", i)
		_, e := fmt.Scanf("%d", &num)
		if e != nil {
			fmt.Println(e)
			fmt.Println("Exiting ...")
			return
		}
		s = append(s, num)
	}
	BubbleSort(s)
	fmt.Printf("The sorted slice is : %v\n", s)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + t*v0 + s0
	}
	return  fn
}
func main() {
	var a, v0, s0 float64
	var err error
	fmt.Printf("Enter the acceleration : ")
	_, err = fmt.Scanf("%f", &a)
	if err != nil{
		fmt.Println(err)
		fmt.Printf("Exiting ...")
		return
	}
	fmt.Printf("Enter the initial velocity : ")
	_, err = fmt.Scanf("%f", &v0)
	if err != nil{
		fmt.Println(err)
		fmt.Printf("Exiting ...")
		return
	}
	fmt.Printf("Enter the initial displacement : ")
	_, err = fmt.Scanf("%f", &s0)
	if err != nil{
		fmt.Println(err)
		fmt.Printf("Exiting ...")
		return
	}

	Displacement := GenDisplaceFn(a, v0, s0)
	var timestr string
	for{
		fmt.Println("Enter the time (or Enter X to terminate) : ")
		fmt.Scan(&timestr)
		if(strings.HasPrefix(timestr, "X")){
			fmt.Println("Exiting ...")
			break
		}
		t, err := strconv.ParseFloat(timestr, 64)
		if err!= nil{
			fmt.Println(err)
			fmt.Println("Exiting ...")
			break
		}
		fmt.Printf("The displacement after %f seconds is %f", t, Displacement(t))
		fmt.Println()
	}

}

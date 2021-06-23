package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name : ")
	name, _ := consoleReader.ReadString('\n')
	fmt.Println("Enter Address : ")
	address, _ := consoleReader.ReadString('\n')


	var m = map[string]string{
		"name": name,
		"address": address,
	}
	j, _ := json.Marshal(m)
	fmt.Println(string(j))
}

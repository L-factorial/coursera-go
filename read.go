package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Person struct {
	Firstname, Lastname string
}
func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	filename, _ := consoleReader.ReadString('\n')
	fmt.Println(filename)
	filename = strings.TrimSpace(filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Invalid file")
	}
	scanner := bufio.NewScanner(file)

	persons := make([]Person, 0)
	for scanner.Scan() {
		line := scanner.Text()

		linescanner := bufio.NewScanner(strings.NewReader(line))


		linescanner.Scan()
		firstname := linescanner.Text()
		linescanner.Scan()
		lastname := linescanner.Text()

		fmt.Println(firstname)
		fmt.Println(lastname)

		person := Person{firstname, lastname}
		persons = append(persons, person)

	}
	fmt.Printf("%v\n", persons)
}

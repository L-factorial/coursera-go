package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/
func main() {
	var str string
	fmt.Println("Enter the string :")

	consoleReader := bufio.NewReader(os.Stdin)

	str, _ = consoleReader.ReadString('\n')

	var found bool
	s1 := strings.ToLower(str)
	fmt.Println(s1)
	suffix := "n\n"
	cond1 := strings.IndexRune(s1, 'i') == 0
	cond2 := strings.LastIndexAny(s1, suffix) == len(s1) - 1
	cond3 := strings.ContainsRune(s1, 'a')

	found = cond1 &&  cond2 && cond3
	if found {
		fmt.Println("Found")
	}

}

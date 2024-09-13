package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	textfile, _ := os.Open("input.txt")
	scanner1 := bufio.NewScanner(textfile)
	sum := 0
	var fileLines string
	for scanner1.Scan() {
		fileLines = scanner1.Text()
		if ValidString(fileLines) {
			sum++
		}
	}
	fmt.Println(sum)
}

func ValidString(s string) bool {
	if check1(s) && check2(s) {
		return true
	}
	return false
}

func check1(s string) bool {
	sub := ""
	sub1 := ""
	for i := 0; i < len(s)-1; i++ {
		sub = s[i : i+2]
		for j := i + 2; j < len(s)-1; j++ {
			sub1 = s[j : j+2]
			fmt.Println(sub, sub1)
			if sub == sub1 {
				return true
			}
		}
	}
	return false
}

func check2(s string) bool {
	for i := 0; i < len(s); i++ {
		if i < len(s)-2 && s[i] == s[i+2] {
			return true
		}
	}
	return false
}

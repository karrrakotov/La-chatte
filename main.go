package main

import "fmt"

func main() {
	isValid("{}()[][]()")
	isValid("{(})")
	isValid("{((({})))")
}

func isValid(str string) {
	if len(str)%2 != 0 {
		fmt.Println("Скобки несбалансированы")
		return
	}

	var openBrackets []byte
	closedBrackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for idx := range str {
		if str[idx] == '(' || str[idx] == '[' || str[idx] == '{' {
			openBrackets = append(openBrackets, str[idx])
		} else if str[idx] == ')' || str[idx] == ']' || str[idx] == '}' {
			if len(openBrackets) <= 0 {
				fmt.Println("Скобки несбалансированы")
				return
			}
			if openBrackets[len(openBrackets)-1] == closedBrackets[str[idx]] {
				openBrackets[len(openBrackets)-1] = 0
				openBrackets = openBrackets[:len(openBrackets)-1]
			} else {
				fmt.Println("Скобки несбалансированы")
				return
			}
		} else {
			fmt.Println("incorrect brackets")
			return
		}
	}

	fmt.Println("Скобки сбалансированы")
}

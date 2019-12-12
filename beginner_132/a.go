package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()
	num1 := 1
	num2 := 0
	if s[1] == s[0] {
		num1 += 1
		if s[2] != s[0] {
			num2 += 1
			if s[3] == s[2] {
				fmt.Println("Yes")
				return
			}
		}

	} else {
		num2 += 1
		if s[2] == s[0] {
			num1 += 1
			if s[3] == s[1] {
				fmt.Println("Yes")
				return
			}
		} else if s[2] == s[1] {
			if s[3] == s[0] {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}

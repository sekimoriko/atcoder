package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()
	s := nextLine()
	front := s[0 : k-1]
	back := s[k:n]
	lower := strings.ToLower(string(s[k-1]))
	// s[k] = s[k+1]
	// fmt.Println(front)
	// fmt.Println(back)
	// fmt.Println(lower)
	fmt.Println(front + lower + back)
	// fmt.Println(s[k])
	// fmt.Println(n)
}

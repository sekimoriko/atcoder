package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	sc.Split(bufio.ScanWords)
	h := nextInt()
	w := nextInt()
	a := nextInt()
	b := nextInt()

	s0 := ""
	s1 := ""

	for j := 0; j < w-a; j++ {
		s0 += "1"
		s1 += "0"
	}
	for j := 0; j < a; j++ {
		s0 += "0"
		s1 += "1"
	}

	for i := 0; i < h-b; i++ {
		fmt.Println(s0)
	}
	for i := 0; i < b; i++ {
		fmt.Println(s1)
	}
}

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
	a := nextInt()
	b := nextInt()

	if 1 <= a && a <= 9 && 1 <= b && b <= 9 {
		fmt.Println(a * b)
	} else {
		fmt.Println(-1)
	}
}

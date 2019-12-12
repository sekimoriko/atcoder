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
	m1 := nextInt()
	var _ = nextInt()
	m2 := nextInt()
	var _ = nextInt()

	if m1 != m2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

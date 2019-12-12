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

	num := 0
	port := 1

	for i := 0; i < 100; i++ {
		if port < b {
			port += (a - 1)
			num++
		}
	}
	fmt.Println(num)
}

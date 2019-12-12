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
	n := nextInt()
	d := nextInt()
	if n%(2*d+1) != 0 {
		fmt.Println(n/(2*d+1) + 1)
	} else {
		fmt.Println(n / (2*d + 1))
	}
}

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
	k := nextInt()
	x := nextInt()
	y := nextInt()
	if k % 2 == 0 && (x + y) % 2 == 1 {
		fmt.Println(-1)
		return
	}
	
	fmt.Println(n)
}

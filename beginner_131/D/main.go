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

func scanNumArrays(len_n int, len_m int) (arrays [][]int) {
	arrays = make([][]int, len_n)
	for i := 0; i < len_n; i++ {
		arrays[i] = make([]int, len_m)
	}
	for i := 0; i < len_n; i++ {
		for j := 0; j < len_m; j++ {
			arrays[i][j] = nextInt()
		}
	}
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	ab := scanNumArrays(n, 2)

	


}

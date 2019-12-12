package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	x := nextInt()
	a := make([]int, 6000)

	for i := 0; i < 6; i++ {
		for j := 0; j < 1000; j++ {
			index := i*1000 + j
			a[index] = (100 + i) * (j + 1)
		}
	}
	sort.Ints(a)

	dp := make([][]bool, 6010)
	for i := 0; i < 6010; i++ {
		dp[i] = make([]bool, 100010)
	}
	dp[0][0] = true

	for i := 0; i < 6000; i++ {
		for j := 0; j <= x; j++ {
			if j >= a[i] {
				dp[i+1][j] = dp[i][j] || dp[i][j-a[i]]
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	if dp[6000][x] {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

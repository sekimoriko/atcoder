package main

import (
	"bufio"
	"fmt"
	"math"
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
	d := nextInt()
	x := scanNumArrays(n, d)

	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			distance := 0
			check := 0.
			for k := 0; k < d; k++ {
				distance += (x[i][k] - x[j][k]) * (x[i][k] - x[j][k])
			}
			check = math.Sqrt(float64(distance))
			if math.Floor(check) == check {
				count++
			}
		}
	}
	fmt.Println(count)
}

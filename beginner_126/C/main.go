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

func checkBi(n int, k int) int {
	pow := 0
	div := n
	for {
		if div < k {
			div = div * 2
			pow++
		} else {
			return pow
		}
		// fmt.Println(div)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()
	p := 0.0

	// if n >= k {
	// 	p = float64(k) / float64(n)
	// }

	ptmp := float64(1) / float64(n)
	// fmt.Printf("%.10f\n", ptmp)
	for i := 1; i < n+1; i++ {
		p += ptmp / float64(math.Pow(2.0, float64(checkBi(i, k))))
	}

	fmt.Printf("%.12f\n", p)
}

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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	r := int(math.Floor(math.Sqrt(float64(n))))
	for i := r; i > 0; i-- {
		m := n / i
		d := n % i
		if d == 0 {
			sum := m + i - 2
			fmt.Println(sum)
			return
		}
	}
}

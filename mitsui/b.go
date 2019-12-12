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

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	x := float64(n) / 1.08
	x2 := math.Ceil(x)
	x3 := math.Floor(x)

	n2 := int(x2 * 1.08)
	n3 := int(x3 * 1.08)
	if n == n2 {
		fmt.Println(x2)
	} else if n == n3 {
		fmt.Println(x3)
	} else {
		fmt.Println(":(")
	}
}

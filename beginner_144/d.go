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
	a := float64(nextInt())
	b := float64(nextInt())
	x := float64(nextInt())

	var tan float64
	if x > a*a*b/2 {
		tan = 2 * (a*a*b - x) / (a * a * a)
	} else {
		tan = a * b * b / (2 * x)
	}
	theta := math.Atan(tan)
	c := theta * 180 / math.Pi
	fmt.Println(c)
}

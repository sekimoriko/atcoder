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

func gcd(a, b int) int {
    if b == 0 {
        return a
    } else {
        return gcd(b, a % b)
    }
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()
	numC := 0
	numD := 0
	numCAndD := 0

	minC := a / c
	maxC := b / c
	minCRem := a % c

	minD := a / d
	maxD := b / d
	minDRem := a % d
	lcmCD := lcm(c, d)

	minCAndD := a / lcmCD
	maxCAndD := b / lcmCD
	minCAndDRem := a % lcmCD

	if minCRem != 0 {
		numC = maxC - minC
	} else {
		numC = maxC - minC + 1
	}
	if minDRem != 0 {
		numD = maxD - minD
	} else {
		numD = maxD - minD + 1
	}
	if minCAndDRem != 0 {
		numCAndD = maxCAndD - minCAndD
	} else {
		numCAndD = maxCAndD - minCAndD + 1
	}
	fmt.Println(b - a + 1 - (numC + numD - numCAndD))
}

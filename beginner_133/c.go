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
	l := nextInt()
	r := nextInt()
	if r-l < 2019 {
		l = l % 2019
		r = r % 2019
	} else {
		l = l % 2019
		r = 2019
	}

	num1 := 2019

	for i := l; i < r; i++ {
		for j := i + 1; j < r+1; j++ {
			if i*j%2019 < num1 {
				num1 = i * j % 2019
			}
		}
	}

	fmt.Println(num1)
}

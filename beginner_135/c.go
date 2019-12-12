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

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		num = nextInt()
		nums = append(nums, num)
	}
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := scanNums(n + 1)
	b := scanNums(n)
	monster := 0
	for i := 0; i < n; i++ {
		for j := i; j < i+2; j++ {
			if b[i] >= a[j] {
				b[i] -= a[j]
				monster += a[j]
				a[j] = 0
			} else {
				a[j] -= b[i]
				monster += b[i]
				b[i] = 0
			}
		}
	}
	fmt.Println(monster)
}

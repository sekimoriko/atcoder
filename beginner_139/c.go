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

func max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	h := scanNums(n)

	num := make([]int, n)
	for i := range num {
		num[i] = 1
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			if h[i] <= h[i-1] {
				num[i] += num[i-1]
			} else {
				num[i] = 1
			}
		}
	}
	fmt.Println(max(num) - 1)
}

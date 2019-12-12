package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
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
	d := scanNums(n)
	sort.Ints(d)
	if d[(n/2) - 1] == d[n/2] {
		fmt.Println(0)
	} else {
		fmt.Println(d[n/2] - d[(n/2) - 1])
	}
}

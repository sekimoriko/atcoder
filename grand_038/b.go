package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
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
	k := nextInt()
	p := scanNums(n)

	count := 1
	for i := 0; i < n-k+1; i++ {
		p2 := make([]int, k)
		copy(p2, p[i:i+k])
		sort.Ints(p2)
		// fmt.Print(p[i : i+k])
		// fmt.Println(p2)
		if !reflect.DeepEqual(p[i:i+k], p2) {
			count++
		}
	}
	fmt.Println(count)
}

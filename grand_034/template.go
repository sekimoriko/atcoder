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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		num = nextInt()
		nums = append(nums, num)
	}
	return
}

func scanNumArrays(lenN int, lenM int) (arrays [][]int) {
	arrays = make([][]int, lenN)
	for i := 0; i < lenN; i++ {
		arrays[i] = make([]int, lenM)
	}
	for i := 0; i < lenN; i++ {
		for j := 0; j < lenM; j++ {
			arrays[i][j] = nextInt()
		}
	}
	return
}

func scanStrings(len int) (strings []string) {
	var str string
	for i := 0; i < len; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	d := scanNumArrays(n, m)
	fmt.Println(d)
	// n := nextInt()
	//
	// graph := make([][]string, n)
	// for i := 0; i < n; i++ {
	// 	graph[i] = make([]string, m)
	// }
}

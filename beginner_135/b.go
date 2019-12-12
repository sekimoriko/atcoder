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

func checkUpper(nums []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	p := scanNums(n)
	upper := false

	if checkUpper(p, n) {
		fmt.Println("YES")
		return
	}

	for i := 0; i < n-1; i++ {
		if p[i] > p[i+1] {
			for j := 0; j < n; j++ {
				var tmp int
				if i != j {
					p2 := make([]int, n)
					copy(p2, p)
					tmp = p2[i]
					p2[i] = p2[j]
					p2[j] = tmp
					// fmt.Println(p2)
					upper = checkUpper(p2, n)
					if upper {
						fmt.Println("YES")
						return
					}
				}
			}
		}
	}
	fmt.Println("NO")
}

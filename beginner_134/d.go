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

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		num = nextInt()
		nums = append(nums, num)
	}
	return
}

func refbit(i int, b int) int {
	return (i >> uint8(b)) & 1
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := scanNums(n)
	patternNum := int(math.Pow(2, float64(n)))
	count := 0
	correctCount := 0
	correctBox := make([]bool, n)
	for i := 0; i < patternNum; i++ {
		box := make([]bool, n)
		count = 0
		for j := 0; j < n; j++ {
			if refbit(i, j) == 1 {
				box[j] = true
			} else {
				box[j] = false
			}
		}
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				if (k+1)%(j+1) == 0 {
					if box[k] == true {
						sum++
					}
				}
			}
			if sum%2 == a[j] {
				count++
			}
		}
		if count == n {
			correctCount++
			correctBox = box
		}
	}

	if correctCount == 1 {
		falseCount := 0
		for i := 0; i < n; i++ {
			if correctBox[i] == false {
				falseCount++
			}
		}
		if falseCount == n {
			fmt.Println(0)
			return
		}
	}

	if correctCount > 0 {
		num := 0
		for i := 0; i < n; i++ {
			if correctBox[i] {
				num++
			}
		}
		fmt.Println(num)
		num = 0
		for i := 0; i < n; i++ {
			if correctBox[i] {
				if num > 0 {
					fmt.Print(" ")
				}
				fmt.Print(i + 1)
				num++
			}
		}
		// fmt.Println()
	} else {
		fmt.Println(-1)
	}
}

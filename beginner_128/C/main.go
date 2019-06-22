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

func refbit(i int, b int) int {
	return (i >> uint8(b)) & 1
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	k := make([]int, m)
	s := make([][]int, m)
	p := make([]int, m)

	for i := 0; i < m; i++ {
		k[i] = nextInt()
		s[i] = make([]int, k[i])
		for j := 0; j < k[i]; j++ {
			s[i][j] = nextInt()
		}
	}
	for i := 0; i < m; i++ {
		p[i] = nextInt()
	}

	patternNum := int(math.Pow(2, float64(n)))
	// fmt.Println(patternNum)
	switches := make([]bool, n)
	nums := 0

	for i := 0; i < patternNum; i++ {
		for j := 0; j < n; j++ {
			// fmt.Println(refbit(patternNum, j))
			if refbit(i, j) == 1 {
				switches[j] = true
			} else {
				switches[j] = false
			}
		}

		onLightNum := 0
		for light := 0; light < m; light++ {
			count := 0
			for sw := 0; sw < k[light]; sw++ {
				if switches[s[light][sw]-1] {
					count++
				}
			}
			if (count % 2) == p[light] {
				onLightNum++
			}
		}
		if onLightNum == m {
			nums++
		}
	}
	fmt.Println(nums)

	// fmt.Println(n)
	// fmt.Println(m)
	// fmt.Println(k)
	// fmt.Println(s)
	// fmt.Println(p)
}

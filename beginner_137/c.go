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

func scanStrings(len int) (strings []string) {
	for i := 0; i < len; i++ {
		sc.Scan()
		strings = append(strings, sc.Text())
	}
	return
}

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := scanStrings(n)

	word := make([][]int, n)
	for i := 0; i < n; i++ {
		word[i] = make([]int, 26)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			word[i][j] = 0
		}
	}
	// fmt.Println(getRuneAt(s[1], 1))
	for i := 0; i < n; i++ {
		// for j := 0; j < 10; j++ {
		for _, j := range s[i] {
			// getRuneAt(s[i], j)
			// word[i][getRuneAt(s[i], j)-[]rune("a")[0]]++
			// fmt.Print(j)
			word[i][j-97]++
		}
	}
	// fmt.Println(s[0])

	wordi := make([]int64, 1000000000)
	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			wordi[i] += int64(word[i][j]) * int64(math.Pow10(j))
		}
	}

	count := int64(0)

	// for i := 0; i < n-1; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		if reflect.DeepEqual(word[i], word[j]) {
	// 			count++
	// 		}
	// 	}
	// }

	encount := map[int64]int64{}

	for i := 0; i < len(s); i++ {
		if encount[wordi[i]] == 0 {
			encount[wordi[i]] = 1
		} else {
			count += encount[wordi[i]]
			encount[wordi[i]]++
		}
	}
	fmt.Println(count)
}

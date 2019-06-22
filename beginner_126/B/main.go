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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func checkMM(mm string) bool {
	mmI, _ := strconv.Atoi(mm)
	if (mmI == 0) || (mmI >= 13) {
		return false
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextLine()
	front := string(n[0:2])
	back := string(n[2:4])
	checkFront := checkMM(front)
	checkBack := checkMM(back)

	if checkFront && checkBack {
		fmt.Println("AMBIGUOUS")
	} else if checkFront && !checkBack {
		fmt.Println("MMYY")
	} else if !checkFront && checkBack {
		fmt.Println("YYMM")
	} else {
		fmt.Println("NA")
	}
}

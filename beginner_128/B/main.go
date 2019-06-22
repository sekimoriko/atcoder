package main

import (
	"bufio"
	"fmt"
	"os"
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

type restaurant struct {
	index int
	city  string
	point int
}

type lessFunc func(i, j *restaurant) bool

type sorter struct {
	res      []restaurant
	lessfunc []lessFunc
}

func (is sorter) Len() int {
	return len(is.res)
}
func (is sorter) Swap(i, j int) {
	is.res[i], is.res[j] = is.res[j], is.res[i]
}
func (is sorter) Less(i, j int) bool {
	var k int
	p, q := &is.res[i], &is.res[j]
	for k = 0; k < len(is.lessfunc)-1; k++ {
		less := is.lessfunc[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return is.lessfunc[k](p, q)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	r := make([]restaurant, n)
	for i := 0; i < n; i++ {
		r[i].index = i + 1
		r[i].city = nextLine()
		r[i].point = nextInt()
	}
	byCity := func(p1, p2 *restaurant) bool {
		return p1.city < p2.city
	}
	byPoint := func(p1, p2 *restaurant) bool {
		return p1.point > p2.point
	}
	sort.Sort(sorter{res: r, lessfunc: []lessFunc{byCity, byPoint}})
	// sort.SliceStable(r, func(i int, j int) bool {
	// 	p, q := r[i], r[j]
	// 	if p.city < q.city {
	// 		return true
	// 	} else if p.city > q.city {
	// 		return false
	// 	} else {
	// 		return p.point > q.point
	// 	}
	// })
	for _, v := range r {
		fmt.Println(v.index)
	}
}

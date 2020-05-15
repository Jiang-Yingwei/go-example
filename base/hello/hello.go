package hello

import (
	"fmt"
	"time"
)

func RangeTest() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go run" {
		fmt.Println(i, c)
	}
}

func MapTest() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))
	fmt.Println("v1:", m["k1"])

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs1 := m["k1"]
	_, prs2 := m["k2"]
	fmt.Println("prs:", prs1, prs2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func SliceTest() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("init len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	//左边开区间，右边闭区间
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func ArrayTest() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set index 4", a)
	fmt.Println("get index 4", a[4])
	fmt.Println("length:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	var twoD [2][3][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				twoD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("2d:", twoD, len(twoD))

}

func SwitchTest() {
	i := 2
	fmt.Println(i, "as")
	switch i {
	case 2:
		fmt.Println("two")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("today is the weekend")
	default:
		fmt.Println("today is a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("now is before noon")
	default:
		fmt.Println("now is after noon")
	}
}

func ForTest() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}

func IfTest(num int) {
	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	if a := 9; num > a {
		fmt.Println("num is bigger then 9")
	}
}

func varTest()  {
	var a string = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}
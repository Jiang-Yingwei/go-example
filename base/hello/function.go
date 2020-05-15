package hello

import (
	"fmt"
	"math"
	"time"
)

func ValuesAndParams(nums ...int) (int, int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return len(nums), total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func ClosuresTest() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt2 := intSeq()
	fmt.Println(newInt2())

	s1 := []string{"a", "b", "c"}
	for _, v := range s1 {
		go func() {
			fmt.Println("s1", v)
		}()
	}

	s2 := []string{"a", "b", "c"}
	for _, v := range s2 {
		go func(v string) {
			fmt.Println("s2", v)
		}(v) //每次将变量 v 的拷贝传进函数
	}
	time.Sleep(time.Second * 1)
}

func zeroVal(iVal int) {
	iVal = 0
}
func zeroPtr(iPtr *int) {
	*iPtr = 0
}
func PointersTest() {
	i := 1
	fmt.Println("initial:", i)
	//zeroval 有一个 int 型参数，所以使用值传递。zeroval 将从调用它的那个函数中得到一个 ival形参的拷贝。
	zeroVal(i)
	fmt.Println("zeroval:", i)
	fmt.Println("pointer:", &i)
	//函数体内的 *iptr 接着解引用 这个指针，从它内存地址得到这个地址对应的当前值。对一个解引用的指针赋值将会改变这个指针引用的真实地址的值
	zeroPtr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}

type person struct {
	name string
	age  int
}

func (p *person) older() {
	p.age++
}

func (p person) younger() int {
	//未带指针不影响原结构体
	p.age--
	return p.age
}

func StructTest() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	//取地址
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
	sp.older()
	fmt.Println(s.age)
	fmt.Println(sp.younger())
	fmt.Println(s.age)
	fmt.Println(sp.age)
}

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func InterfaceTest() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

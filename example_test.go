package slice_test

import (
	"fmt"

	"github.com/sina-devel/slice"
)

func ExampleEqual() {
	s1 := []int8{20, 3, 4}
	s2 := []int8{20, 3, 4}
	fmt.Println(slice.Equal(s1, s2))
	//Output:
	// true
}

func ExampleEqualFunc() {
	s1 := []int8{20, 2}
	s2 := []int16{20, 2}

	fmt.Println(slice.EqualFunc(s1, s2, func(a int8, b int16) bool {
		return int16(a) == b
	}))
	//Output:
	// true
}

func ExampleIndex() {
	s := []int{10, 20, 40}
	fmt.Println("20 exists at index", slice.Index(s, 20))
	//Output:
	// 20 exists at index 1
}

func ExampleIndexFunc() {
	type Point struct {
		x, y int16
	}

	points := []Point{{10, 2}, {4, 4}}

	fmt.Println("point(10, _) exists at index", slice.IndexFunc(points, func(p Point) bool {
		return p.x == 10
	}))
	//Output:
	// point(10, _) exists at index 0
}

func ExampleContains() {
	animals := []string{"zebra", "lion", "gopher"}

	if slice.Contains(animals, "gopher") {
		fmt.Println("ʕ◔ϖ◔ʔ")
	}
	//Output:
	// ʕ◔ϖ◔ʔ
}

func ExampleClone() {
	numbers := []float64{1.0, 10.2, 39.2}

	cloned := slice.Clone(numbers)
	cloned[0] = 23.3

	fmt.Println(numbers, cloned)
	//Output:
	// [1 10.2 39.2] [23.3 10.2 39.2]
}

func ExampleSort() {
	s := []rune{'🥚', '🐔'}
	slice.Sort(s, func(a, b rune) bool { return a < b })
	fmt.Printf("%q", s)
	//Output:
	// ['🐔' '🥚']
}

package link

import "fmt"

func BuildList(limit int) *ConsCell {
	list := NIL
	for i := 0; i < limit; i++ {
		list = list.Cons(i)
	}
	return list
}

func ExampleConsCell_Cons() {
	x := BuildList(5)
	fmt.Println(x)
	// Output:
	// (4, 3, 2, 1, 0)
}

func ExampleConsCell_Len() {
	x := BuildList(5)
	fmt.Println(x.Len())
	fmt.Println(NIL.Len())
	// Output:
	// 5
	// 0
}

func ExampleConsCell_Reverse() {
	x := BuildList(5)
	fmt.Println(x.Reverse())
	// Output:
	// (0, 1, 2, 3, 4)
}

func ExampleConsCell_ReverseFirstN() {
	x := BuildList(5)
	fmt.Println(x.Reverse().ReverseFirstN(3))
	// Output:
	// (2, 1, 0) (0) (3, 4)
}

func ExampleConsCell_ReverseBy3s() {
	x := BuildList(20)
	fmt.Println(x.ReverseBy3s())
	// Output:
	// (17, 18, 19, 14, 15, 16, 11, 12, 13, 8, 9, 10, 5, 6, 7, 2, 3, 4, 0, 1)
}

package equal

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	one, oneAgain, two := 1, 1, 2

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1, ch2 := make(chan int), make(chan int)
	var chiro <- chan int = ch1

	type mystring string

	var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

	for _, test := range []struct {
		x, y interface{}
		want bool
	} {
		// basic types
		{1, 1, true},
		{ch1, ch2, true},
		{chiro, ch1, false},
		{&iface1, &iface2, false},
		{iface1Again, &iface1, true},
		{mystring("foo"), "foo", false},		// different types
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("equal %v, %v = %t",
				test.x, test.y, !test.want)
		}
	}
}

func Example_equal() {

	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))        // "true"
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))      // "false"
	fmt.Println(Equal([]string(nil), []string{}))             // "true"
	fmt.Println(Equal(map[string]int(nil), map[string]int{})) // "true"

	// Output:
	// true
	// false
	// true
	// true
}

func Example_equalCycle() {

	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a)) // "true"
	fmt.Println(Equal(b, b)) // "true"
	fmt.Println(Equal(c, c)) // "true"
	fmt.Println(Equal(a, b)) // "false"
	fmt.Println(Equal(a, c)) // "false"

	// Output:
	// true
	// true
	// true
	// false
	// false
}
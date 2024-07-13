package main

import (
	"by-example/lib"
	"fmt"
)

func main() {
	fmt.Println(lib.SayHello())
	fmt.Println(lib.Greet())
	fmt.Printf("%v plus %v is %v\n", 3, 12, lib.Add(3, 12))
	fmt.Printf("%v divided by %v is %v\n", 3.0, 12.0, lib.Div(3, 12))
	fmt.Printf("%v equals %v : %v\n", 3.0, 12.0, lib.Equal(3, 12))
	swap1, swap2 := lib.Swap(3, 12)
	fmt.Printf("%v, %v swapped is %v, %v\n", 3.0, 12.0, swap1, swap2)
	fmt.Println(lib.LANG)
	fmt.Println(lib.OddNumbers(10))
	fmt.Println(lib.Reverse("afternoon"))
	fmt.Println(lib.Runtime())
	fmt.Println(lib.GetType(true))
	fmt.Println(lib.GetType([2]int{}))

	lib.GoByExamples()

	lib.Slices()

	lib.MapWork()

	fmt.Println(lib.StringRunes("pastor"))

	fmt.Println(lib.Mean(1.5, 1.9))

	inc := lib.Increment(10)
	for x := range 5 {
		fmt.Println(x, inc())
	}

	fmt.Println(lib.Factorial(5))

	i := 1516
	lib.Zeroval(i)
	fmt.Println("i after Zeroval", i)
	lib.Zeroptr(&i)
	fmt.Println("i after Zeroptr", i)

	feedback := lib.Feedback{Comment: "nice", Rating: 3}
	fmt.Println(feedback)
	lib.SetRating(feedback, 1)
	fmt.Println("passing struct as value", feedback)
	lib.SetRatingPtr(&feedback, 1)
	fmt.Println("passing struct as pointer", feedback)
	fmt.Println("timezone", feedback.Timezone())
	fmt.Println(feedback.RatingDescription())

	newFeedback := lib.NewFeedback(4)
	fmt.Printf("%+v\n", newFeedback)
	// pointer receiver methods can be called with a pointer to a struct
	fmt.Println(newFeedback.Timezone(), "timezone")
	// value receiver methods can also be called with a pointer to a struct
	fmt.Println(newFeedback.RatingDescription(), "timezone")

	fmt.Println(lib.CountRunes("halleluyah", 'l'))
	fmt.Println(lib.CountRunes("halleluyah", 'A'))

	// an interface method impl with pointer receiver requires a pointer
	lib.Measure(&lib.Circle{Radius: 2.06})
	//an interface method impl with value receiver can take a value or pointer
	lib.Measure(&lib.Rect{
		Len:   2.2,
		Width: 13.5,
	})
	// an uninitialised struct is not nil, its fields have zero values
	var x lib.Rect
	lib.Measure(x)

	t := 8
	a := []int{1, 2, 4, 9}
	fmt.Printf("the index of %v in %v is %v\n", t, a, lib.Index(a, t))
	c := 'e'
	chars := []rune{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("the index of %v in %v is %v\n", c, chars, lib.Index(chars, c))

	fmt.Println(lib.MapKeys(map[string]int{"one": 1, "two": 2}))

	var lst lib.List[string]
	lst.Add("thank")
	lst.Add("you")
	fmt.Println(lst.ToSlice())

	num := 2
	res, e := lib.Add3(num)
	if e != nil {
		fmt.Printf("Could not add 3 to %v: %v\n", num, e)
	} else {
		fmt.Printf("Adding 3 to %v: %v\n", num, res)
	}
	err := lib.MyError{Vars: map[string]string{"one": "1", "two": "2"}, Key: "org.harmony.test"}
	fmt.Println(err)
}

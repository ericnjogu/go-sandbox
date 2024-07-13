package lib

import (
	"fmt"
	"strings"
)

func Slices() {
	var s []string
	fmt.Println("s", s, "len", len(s))

	var sbm []string = make([]string, 3)
	fmt.Println("sbm", sbm, "len", len(sbm), "cap", cap(sbm), "index 0", "'"+sbm[0]+"'")
	sbm[0] = "s"
	sbm[1] = "b"
	sbm[2] = "m"
	sbm = append(sbm, "x")
	fmt.Println("sbm after append", sbm)

	fmt.Println("slice to 2", sbm[:2])
	fmt.Println("slice from 3", sbm[3:])

	slit := []string{"one", "two", "three"}
	for n := range len(slit) {
		fmt.Printf("%v,", strings.ToUpper(slit[n]))
	}
	fmt.Println()

	vowels := [...]string{"a", "e", "i", "o", "u"}
	arraySlice := vowels[3:]
	appendedSlice := append(arraySlice, "r", "s")
	fmt.Println("appended slice", appendedSlice)
	fmt.Println("vowels", vowels)

	vowels[3] = "x"
	fmt.Println("setting element at index 3 to x")
	fmt.Println("appended slice", appendedSlice)
	fmt.Println("array slice", arraySlice)
	fmt.Println("vowels", vowels)

	for i, v := range vowels {
		fmt.Print(i, "=", v, ",")
	}
	fmt.Println()

	for _, v := range appendedSlice {
		fmt.Print(v, ",")
	}
	fmt.Println()

}

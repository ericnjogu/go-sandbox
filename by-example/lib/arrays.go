package lib

import "fmt"

func GoByExamples() {
	var a [5]int
	fmt.Println("a", a, "array")

	a[4] = 1706
	fmt.Println(a, "after set")
	fmt.Println("len a:", len(a))

	b := [...]int{1, 7, 2, 29, 24}
	fmt.Println("b", b)

	c := [...]int{1, 2, 5: 10}
	fmt.Println("c", c)

	var matrix [3][3]string
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = fmt.Sprintf("%v,%v", i, j)
		}
	}
	fmt.Println(matrix)
}

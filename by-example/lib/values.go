package lib

const LANG string = "en"

var name = "xyz"

func Greet() string {
	return "good" + " " + "morning"
}

func Add(num1, num2 int) int {
	return num1 + num2
}

func Div(num1, num2 float32) float32 {
	return num1 / num2
}

func Equal(num1, num2 float32) bool {
	return num1 == num2
}

func Swap(num1, num2 int) (int, int) {
	return num2, num1
}

func StringRunes(str string) map[int]string {
	m := make(map[int]string)
	for i, c := range str {
		m[i] = string(c)
	}
	return m
}

func CountRunes(source string, target rune) int {
	count := 0
	for _, char := range source {
		if char == target {
			count++
		}
	}
	return count
}

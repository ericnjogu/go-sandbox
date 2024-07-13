package lib

import (
	"fmt"
	"maps"
)

func MapWork() {
	runs := make(map[string]float32)
	runs["juzi"] = 8.3
	runs["leo"] = 11.16
	fmt.Println(runs, len(runs))
	for k, v := range runs {
		fmt.Print(k, "=", v, ",")
	}
	fmt.Println()

	dist, present := runs["kesho"]
	fmt.Println(present, dist)

	delete(runs, "wapi") // no error
	clear(runs)
	fmt.Println("after clearing", runs, len(runs))

	fmt.Println("literal maps are equal:",
		maps.Equal(map[string]int{"one": 1}, map[string]int{"two": 2}))

	type Vertex struct {
		Lat, Long float64
	}
	var pins map[string]Vertex
	pins = make(map[string]Vertex)
	pins["home"] = Vertex{
		Lat:  .908767,
		Long: 19.0979752,
	}
	fmt.Println("pins", pins)
}

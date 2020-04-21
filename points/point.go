package points

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Point struct {
	X, Y int
}

func (p1 Point) Distance(p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func FromJSON(file string) []Point {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var result []Point
	json.Unmarshal(content, &result)

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

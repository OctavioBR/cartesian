package points

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Point struct {
	X, Y int
}

func (p Point) Distance(other Point) int {
	return 0 // distance(p1,p2) = |x1-x2| + |y1-y2|
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

package point

type Point struct {
	X, Y int
}

type PointDistance struct {
	Point
	Distance int
}

func (p1 Point) Distance(p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

package points

import "testing"

func TestDistance(t *testing.T) {
	cases := []struct {
		p1   Point
		p2   Point
		want int
	}{
		{Point{0, 0}, Point{2, 3}, 5},
		{Point{-10, 0}, Point{10, 0}, 20},
		{Point{10, 0}, Point{-10, 0}, 20},
		{Point{0, 10}, Point{0, -10}, 20},
		{Point{-1, 1}, Point{1, -1}, 4},
	}

	for _, c := range cases {
		got := c.p1.Distance(c.p2)

		if got != c.want {
			t.Errorf("Distance between %+v and %+v should be %d, but got %d",
				c.p1, c.p2, c.want, got)
		}
	}
}

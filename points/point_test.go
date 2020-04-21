package points

import "testing"

func TestDistance(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{2, 3}

	want := 5
	got := p1.Distance(p2)

	if got != want {
		t.Errorf("Distance between %+v and %+v should be %d, but got %d",
			p1, p2, want, got)
	}
}

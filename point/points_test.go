package point

import (
	"testing"
)

func TestSortedInsert(t *testing.T) {
	expected := 3

	list := PointsDistance{
		PointDistance{Point{}, 1},
		PointDistance{Point{}, 2},
		PointDistance{Point{}, 4},
		PointDistance{Point{}, 5},
	}

	list = list.SortedInsert(PointDistance{Point{}, expected})
	if list[2].Distance != expected {
		t.Errorf("Distance of list[2] should be %v, got %v", expected, list[2].Distance)
	}
}

func TestSortedInsertEmpty(t *testing.T) {
	list := PointsDistance{}

	list = list.SortedInsert(PointDistance{Point{}, 1})

	if len(list) != 1 {
		t.Error("Inserting in empty PointsDistance should work")
	}
}

func TestSortedInsertFirst(t *testing.T) {
	expected := 0

	list := PointsDistance{
		PointDistance{Point{}, 1},
		PointDistance{Point{}, 2},
		PointDistance{Point{}, 4},
		PointDistance{Point{}, 5},
	}

	list = list.SortedInsert(PointDistance{Point{}, expected})
	if list[0].Distance != expected {
		t.Errorf("Distance of list[0] should be %v, got %v", expected, list[0].Distance)
	}
}

func TestSortedInsertLast(t *testing.T) {
	expected := 6

	list := PointsDistance{
		PointDistance{Point{}, 1},
		PointDistance{Point{}, 2},
		PointDistance{Point{}, 4},
		PointDistance{Point{}, 5},
	}

	list = list.SortedInsert(PointDistance{Point{}, expected})
	if list[4].Distance != expected {
		t.Errorf("Distance of list[4] should be %v, got %v", expected, list[4].Distance)
	}
}

func TestSortedInsertOtherElements(t *testing.T) {
	list := PointsDistance{
		PointDistance{Point{}, 0},
		PointDistance{Point{}, 1},
		PointDistance{Point{}, 3},
		PointDistance{Point{}, 4},
	}

	list = list.SortedInsert(PointDistance{Point{}, 2})
	for i, v := range list {
		if i != v.Distance {
			t.Errorf("Distance at index %v, should equal %v", i, v.Distance)
		}
	}
}

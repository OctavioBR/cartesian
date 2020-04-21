package point

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

type Points []Point

type PointsDistance []PointDistance

func (p *Points) FromJSON(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(content, &p)
}

func (pd PointsDistance) SortedInsert(in PointDistance) PointsDistance {
	index := sort.Search(len(pd), func(i int) bool {
		return pd[i].Distance > in.Distance
	})

	pd = append(pd, PointDistance{})
	copy(pd[index+1:], pd[index:])
	pd[index] = in

	return pd
}

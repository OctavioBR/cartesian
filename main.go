package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/OctavioBR/cartesian/point"
	"github.com/OctavioBR/cartesian/utils"
)

var data point.Points

func main() {
	data.FromJSON("data/points.json")
	log.Printf("Loaded %d points from data/points.json", len(data))
	fmt.Println(data)

	http.HandleFunc("/api/points", handler)

	port := utils.Getenv("PORT", "8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	refPoint, refDistance, err := validateInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var pds point.PointsDistance
	for _, pt := range data {
		distance := refPoint.Distance(pt)
		if distance <= refDistance {
			pds = pds.SortedInsert(point.PointDistance{pt, distance})
		}
	}

	j, err := json.Marshal(pds)
	if err != nil {
		http.Error(w, "Failed to marshal result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(pds) == 0 {
		w.Write([]byte("[]"))
		return
	}

	w.Write(j)
}

func validateInput(r *http.Request) (point.Point, int, error) {
	q := r.URL.Query()
	var err error

	params := map[string]int{
		"x":        0,
		"y":        0,
		"distance": 0,
	}

	for k := range params {
		v := q.Get(k)
		if v == "" {
			err = fmt.Errorf("param %q is required", k)
			return point.Point{}, 0, err
		}
		vn, err := strconv.Atoi(v)
		if err != nil {
			err = fmt.Errorf("%q must be an integer, got %q", k, v)
			return point.Point{}, 0, err
		}
		params[k] = vn
	}

	return point.Point{params["x"], params["y"]}, params["distance"], nil
}

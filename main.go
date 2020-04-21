package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/OctavioBR/cartesian/points"
	"github.com/OctavioBR/cartesian/utils"
)

var data []points.Point

func main() {
	data = points.FromJSON("data/points.json")
	log.Printf("Loaded %d points from data/points.json", len(data))

	http.HandleFunc("/api/points", handler)

	port := utils.Getenv("PORT", "8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	params := map[string]int{
		"x":        0,
		"y":        0,
		"distance": 0,
	}

	for k := range params {
		v := q.Get(k)
		if v == "" {
			http.Error(w, fmt.Sprintf("param %q is required", k), http.StatusBadRequest)
			return
		}
		vn, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, fmt.Sprintf("%q must be an integer, got %q", k, v), http.StatusBadRequest)
			return
		}
		params[k] = vn
	}

	fmt.Fprintf(w, "x=%v y=%v distance=%v", params["x"], params["y"], params["distance"])
}

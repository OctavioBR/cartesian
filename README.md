> An exercise from [tobym/cartesian](https://github.com/tobym/cartesian)

Cartesian API
=============

Create an API server in [go](https://golang.org/). It will deal with a series of points represented as (x,y) coordinates on a simple 2-dimensional plane. Take a look at https://en.wikipedia.org/wiki/Cartesian_coordinate_system if you need a refresher on this concept.

It must have an api route at `/api/points` that accepts a `GET` request with the following parameters, and returns a JSON list of points that are within `distance` from `x,y`, using the Manhattan distance method. The points should be returned in order of increasing distance from the search origin.
- `x` integer (required). This represents the `x` coordinate of the search origin.
- `y` integer (required). This represents the `y` coordinate of the search origin.
- `distance` integer (required). This represents the Manhattan distance; points within `distance` from `x` and `y` are returned, points outside are filtered out.

The Manhattan distance is measured "block-wise", as the distance in blocks between any two points in the plane (e.g. 2 blocks down and 3 blocks over for a total of 5 blocks). It is defined as the sum of the horizontal and vertical distances between points on a grid. Formally, where `p1 = (x1, y1)` and `p2 = (x2, y2)`, `distance(p1,p2) = |x1-x2| + |y1-y2|`.

On startup, the API server should read a list of points from `data/points.json`.

# Solution
Code purely with Go and it's standard library.  
For each request, it iterates points from data "_O(n)_", and stores occurrences (_m_) whithin distance from search origin using sorted insert. Took advantage of [`sort.Search`](https://golang.org/pkg/sort/#Search) function which uses bynary search "_O(log m)_", achieving the end asymptotic complexity of _O(n log m)_.

There may be other data sctructures that could better explore the cartesian coordinate system. Storing points from json file _sorted_, and enabling improved search from origin, limited by given distance.


### Build?
`make build` → This will create a static binary file at `bin/cartesian`.

### Run?
`make run`

> Notes: You can set `PORT` environment variable to define which port the server will listen to.

## TODO
- [ ] "make test", instead of "cd point & go test"
- [ ] dockerize "make run" too

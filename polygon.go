package mfcg

import "encoding/json"

// Polygon is an area within a set of Points.
type Polygon struct {
	Width  float64   `json:"width"`
	Coords [][]Point `json:"coordinates"`
}

// coordsToPolygon returns a Polygon corresponding to the provided coordinate
// data. The data must conform to a 2D slice of Points.
func coordsToPolygon(data []byte) (*Polygon, error) {
	var points [][]Point
	if err := json.Unmarshal(data, &points); err != nil {
		return nil, err
	}

	return &Polygon{
		Coords: points,
	}, nil
}

// coordsToPolygons returns a slice of Polygons each corresponding to the
// provided coordinate data. The data must conform to a 3D slice of Points.
func coordsToPolygons(data []byte) ([]Polygon, error) {
	var points [][][]Point
	if err := json.Unmarshal(data, &points); err != nil {
		return nil, err
	}

	polys := make([]Polygon, len(points))
	for i := range points {
		polys[i] = Polygon{Coords: points[i]}
	}

	return polys, nil
}

// geosToPolygons returns a slice of Polygons each corresponding to the
// provided GeometryCollection data. The data must conform to a slice of MFCG's
// proprietary polygon geometries.
func geosToPolygons(data []byte) ([]Polygon, error) {
	var geos []*json.RawMessage
	if err := json.Unmarshal(data, &geos); err != nil {
		return nil, err
	}

	var polys []Polygon
	for _, g := range geos {
		var p Polygon
		if err := json.Unmarshal(*g, &p); err != nil {
			return nil, err
		}

		polys = append(polys, p)
	}

	return polys, nil
}

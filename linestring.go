package mfcg

import "encoding/json"

// LineString is a path between a set of Points.
type LineString struct {
	Width  float64 `json:"width"`
	Coords []Point `json:"coordinates"`
}

// geosToLineStrings returns a slice of LineStrings each corresponding to the
// provided GeometryCollection data. The data must conform to a slice of MFCG's
// proprietary linestring geometries.
func geosToLineStrings(data []byte) ([]LineString, error) {
	var geos []*json.RawMessage
	if err := json.Unmarshal(data, &geos); err != nil {
		return nil, err
	}

	var lines []LineString
	for _, g := range geos {
		var ln LineString
		if err := json.Unmarshal(*g, &ln); err != nil {
			return nil, err
		}

		lines = append(lines, ln)
	}

	return lines, nil
}

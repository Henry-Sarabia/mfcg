package mfcg

import (
	"encoding/json"
	"errors"
)

// Point contains the coordinates of a point on a cartesian plane.
type Point struct {
	X float64
	Y float64
}

// UnmarshalJSON decodes the X and Y coordinates of a Point.
// The data must conform to a slice of float64's of length 2.
func (p *Point) UnmarshalJSON(data []byte) error {
	var points []interface{}
	if err := json.Unmarshal(data, &points); err != nil {
		return err
	}

	var ok bool
	if p.X, ok = points[0].(float64); !ok {
		return errors.New("expecting float64 for Point's X field")
	}
	if p.Y, ok = points[1].(float64); !ok {
		return errors.New("expecting float64 for Point's Y field")
	}

	return nil
}

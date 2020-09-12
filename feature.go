package mfcg

import (
	"encoding/json"
)

// featureCollection contains the list of a map's features.
type featureCollection struct {
	Type     string     `json:"type"`
	Features []*feature `json:"features"`
}

// feature represents a specific type of map feature (e.g. roads, rivers, buildings).
type feature struct {
	MetaData
	Type        string          `json:"type"`
	ID          string          `json:"id"`
	Coordinates json.RawMessage `json:"coordinates"`
	Geometries  json.RawMessage `json:"geometries"`
}

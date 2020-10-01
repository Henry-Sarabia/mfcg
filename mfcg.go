package mfcg

import (
	"encoding/json"
	"io"
)

// New reads the provided MFCG data from r and returns the corresponding Map.
func New(r io.Reader) (*Map, error) {
	var collect featureCollection
	if err := json.NewDecoder(r).Decode(&collect); err != nil {
		return nil, err
	}

	feats := make(map[string]feature)
	for _, ft := range collect.Features {
		if ft.ID == "" {
			continue
		}
		feats[ft.ID] = *ft
	}

	mp, err := toMap(feats)
	if err != nil {
		return nil, err
	}

	return mp, nil
}

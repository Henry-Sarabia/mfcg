package mfcg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_toMap(t *testing.T) {
	table := map[string]feature{
		IDEarth:     {Coordinates: []byte(`[[[11.1, 11.1]]]`)},
		IDPlanks:    {Geometries: []byte(`[{"coordinates": [[22.2, 22.2]]}]`)},
		IDRivers:    {Geometries: []byte(`[{"coordinates": [[33.3, 33.3]]}]`)},
		IDRoads:     {Geometries: []byte(`[{"coordinates": [[44.4, 44.4]]}]`)},
		IDBuildings: {Coordinates: []byte(`[[[[55.5, 55.5]]]]`)},
		IDFields:    {Coordinates: []byte(`[[[[66.6, 66.6]]]]`)},
		IDGreens:    {Coordinates: []byte(`[[[[77.7, 77.7]]]]`)},
		IDPrisms:    {Coordinates: []byte(`[[[[88.8, 88.8]]]]`)},
		IDSquares:   {Coordinates: []byte(`[[[[99.9, 99.9]]]]`)},
		IDWalls:     {Geometries: []byte(`[{"coordinates": [[[10.10, 10.10]]]}]`)},
		IDWater:     {Coordinates: []byte(`[[[[11.11, 11.11]]]]`)},
		IDValues: {MetaData: MetaData{
			RoadWidth:     12,
			RiverWidth:    13.13,
			TowerRadius:   14.14,
			WallThickness: 15.15,
			Generator:     "foo",
			Version:       "bar",
		}},
	}
	tableMapNoValues := Map{
		Earth:     Polygon{Coords: [][]Point{{{X: 11.1, Y: 11.1}}}},
		Planks:    []LineString{{Coords: []Point{{X: 22.2, Y: 22.2}}}},
		Rivers:    []LineString{{Coords: []Point{{X: 33.3, Y: 33.3}}}},
		Roads:     []LineString{{Coords: []Point{{X: 44.4, Y: 44.4}}}},
		Buildings: []Polygon{{Coords: [][]Point{{{X: 55.5, Y: 55.5}}}}},
		Fields:    []Polygon{{Coords: [][]Point{{{X: 66.6, Y: 66.6}}}}},
		Greens:    []Polygon{{Coords: [][]Point{{{X: 77.7, Y: 77.7}}}}},
		Prisms:    []Polygon{{Coords: [][]Point{{{X: 88.8, Y: 88.8}}}}},
		Squares:   []Polygon{{Coords: [][]Point{{{X: 99.9, Y: 99.9}}}}},
		Walls:     []Polygon{{Coords: [][]Point{{{X: 10.1, Y: 10.1}}}}},
		Water:     []Polygon{{Coords: [][]Point{{{X: 11.11, Y: 11.11}}}}},
	}
	tableMapWithValues := tableMapNoValues
	tableMapWithValues.MetaData = MetaData{
		RoadWidth:     12,
		RiverWidth:    13.13,
		TowerRadius:   14.14,
		WallThickness: 15.15,
		Generator:     "foo",
		Version:       "bar",
	}
	tests := []struct {
		name         string
		replaceKey   string
		replaceValue feature
		want         *Map
		wantErr      bool
	}{
		{
			name:         "Valid features",
			replaceKey:   "",
			replaceValue: feature{},
			want:         &tableMapWithValues,
			wantErr:      false,
		},
		{
			name:         "Invalid Earth",
			replaceKey:   IDEarth,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Planks",
			replaceKey:   IDPlanks,
			replaceValue: feature{Geometries: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Rivers",
			replaceKey:   IDRivers,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Roads",
			replaceKey:   IDRoads,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Buildings",
			replaceKey:   IDBuildings,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Fields",
			replaceKey:   IDFields,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Greens",
			replaceKey:   IDGreens,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Prisms",
			replaceKey:   IDPrisms,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Squares",
			replaceKey:   IDSquares,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Walls",
			replaceKey:   IDWalls,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Invalid Water",
			replaceKey:   IDWater,
			replaceValue: feature{Coordinates: []byte(`["foobar"]`)},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Missing Values",
			replaceKey:   IDValues,
			replaceValue: feature{},
			want:         &tableMapNoValues,
			wantErr:      false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			feats := make(map[string]feature)
			for k, v := range table {
				feats[k] = v
			}
			feats[test.replaceKey] = test.replaceValue

			got, err := toMap(feats)
			if (err != nil) != test.wantErr {
				t.Errorf("got: <%v>, want error: <%v>", err, test.wantErr)
				return
			}

			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

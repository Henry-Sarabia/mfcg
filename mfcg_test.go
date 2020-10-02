package mfcg

import (
	"log"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testFileMap       string = "./test_data/map.json"
	testFileMissingID string = "./test_data/mapMissingID.json"
	testFileInvalid   string = "./test_data/mapInvalidFeature.json"
	testFileEmpty     string = "./test_data/emptyArray.json"
	testFileBlank     string = "./test_data/blank.json"
)

func TestNew(t *testing.T) {
	fullMap := Map{
		Earth: Polygon{
			Coords: [][]Point{
				{
					{X: -387.597, Y: -107.006},
					{X: -392.249, Y: -105.345},
					{X: -435.589, Y: 14.775},
				},
			},
		},
		Planks: []LineString{
			{
				Width: 3.2,
				Coords: []Point{
					{X: 128.119, Y: -82.355},
					{X: 153.725, Y: -101.548},
				},
			},
		},
		Rivers: []LineString{
			{
				Width: 20.079037338457553,
				Coords: []Point{
					{X: 93.579, Y: -128.437},
					{X: 67.981, Y: -102.576},
					{X: 59.021, Y: -60.776},
				},
			},
		},
		Roads: []LineString{
			{
				Width: 8,
				Coords: []Point{
					{X: 0.747, Y: -26.841},
					{X: -59.341, Y: -55.113},
					{X: -131.14, Y: -70.752},
				},
			},
			{
				Width: 8,
				Coords: []Point{
					{X: 26.734, Y: -2.507},
					{X: 56.657, Y: -15.577},
					{X: 120.094, Y: 9.553},
				},
			},
		},
		Buildings: []Polygon{
			{
				Coords: [][]Point{
					{
						{X: 45.922, Y: 13.742},
						{X: 45.238, Y: 5.889},
					},
				},
			},
			{
				Coords: [][]Point{
					{
						{X: 28.763, Y: 15.236},
						{X: 28.483, Y: 27.601},
					},
				},
			},
		},
		Fields: []Polygon{
			{
				Coords: [][]Point{
					{
						{X: -78.086, Y: -37.908},
						{X: -91.414, Y: -40.811},
					},
				},
			},
			{
				Coords: [][]Point{
					{
						{X: -106.891, Y: -86.758},
						{X: -69.06, Y: -78.517},
					},
				},
			},
		},
		Greens: []Polygon{},
		Prisms: []Polygon{
			{
				Coords: [][]Point{
					{
						{X: 12.019, Y: 4.231},
						{X: 9.822, Y: -1.073},
						{X: 4.518, Y: -3.27},
					},
				},
			},
		},
		Squares: []Polygon{
			{
				Coords: [][]Point{
					{
						{X: 26.734, Y: -2.507},
						{X: 0.747, Y: -26.841},
						{X: -26.734, Y: 2.507},
					},
				},
			},
		},
		Walls: []Polygon{
			{
				Width: 7.6,
				Coords: [][]Point{
					{
						{X: -68.846, Y: 46.65},
						{X: -38.751, Y: 75.532},
						{X: 3.733, Y: 87.117},
					},
				},
			},
		},
		Water: []Polygon{
			{
				Coords: [][]Point{
					{
						{X: 93.579, Y: -128.437},
						{X: 139.633, Y: -66.994},
						{X: 201.401, Y: -53.259},
					},
				},
			},
		},
		MetaData: MetaData{
			RoadWidth:     8,
			RiverWidth:    20.079037338457553,
			TowerRadius:   7.6,
			WallThickness: 7.6,
			Generator:     "mfcg",
			Version:       "0.7.7a",
		},
	}
	tests := []struct {
		name    string
		file    string
		want    *Map
		wantErr bool
	}{
		{
			name:    "Valid Reader",
			file:    testFileMap,
			want:    &fullMap,
			wantErr: false,
		},
		{
			name:    "Missing ID field",
			file:    testFileMissingID,
			want:    &Map{},
			wantErr: false,
		},
		{
			name:    "Invalid Feature",
			file:    testFileInvalid,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Empty array",
			file:    testFileEmpty,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "No data",
			file:    testFileBlank,
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				log.Fatal(err)
			}

			got, err := New(f)
			if (err != nil) != test.wantErr {
				t.Errorf("got: <%v>, want error: <%v>", err, test.wantErr)
				return
			}

			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

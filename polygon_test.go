package mfcg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_coordsToPolygon(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Polygon
		wantErr bool
	}{
		{
			name: "Single slice",
			args: args{
				data: []byte(`[[[11.1, 11.1], [22.2, 22.2], [33.3, 33.3]]]`),
			},
			want:    &Polygon{Coords: [][]Point{{{X: 11.1, Y: 11.1}, {X: 22.2, Y: 22.2}, {X: 33.3, Y: 33.3}}}},
			wantErr: false,
		},
		{
			name: "Multiple slices",
			args: args{
				data: []byte(`[
					[[11.1, 11.1]],
					[[22.2, 22.2]],
					[[33.3, 33.3]]
					]`),
			},
			want: &Polygon{Coords: [][]Point{
				{{X: 11.1, Y: 11.1}},
				{{X: 22.2, Y: 22.2}},
				{{X: 33.3, Y: 33.3}},
			}},
		},
		{
			name: "Zero slices",
			args: args{
				data: []byte(`[]`),
			},
			want: &Polygon{
				Coords: [][]Point{},
			},
			wantErr: false,
		},
		{
			name: "No data",
			args: args{
				data: []byte(``),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid data type",
			args: args{
				data: []byte(`[[["foo", "bar"]]]`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := coordsToPolygon(test.args.data)
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

func Test_coordsToPolygons(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []Polygon
		wantErr bool
	}{
		{
			name: "Single 2D slice",
			args: args{
				data: []byte(`[
					[[[11.1, 11.1], [22.2, 22.2], [33.3, 33.3]]]
					]`),
			},
			want:    []Polygon{{Coords: [][]Point{{{X: 11.1, Y: 11.1}, {X: 22.2, Y: 22.2}, {33.3, 33.3}}}}},
			wantErr: false,
		},
		{
			name: "Multiple 2D slices",
			args: args{
				data: []byte(`[
					[[[11.1, 11.1]], [[22.2, 22.2]], [[33.3, 33.3]]]
				]`),
			},
			want:    []Polygon{{Coords: [][]Point{{{X: 11.1, Y: 11.1}}, {{X: 22.2, Y: 22.2}}, {{X: 33.3, Y: 33.3}}}}},
			wantErr: false,
		},
		{
			name: "Zero 2D slices",
			args: args{
				data: []byte(`[]`),
			},
			want:    []Polygon{},
			wantErr: false,
		},
		{
			name: "No data",
			args: args{
				data: []byte(``),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid data type",
			args: args{
				data: []byte(`[[[["foo", "bar"]]]]`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := coordsToPolygons(test.args.data)
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

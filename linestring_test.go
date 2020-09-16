package mfcg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_geosToLineStrings(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    []LineString
		wantErr bool
	}{
		{
			name: "Single geometry",
			args: args{
				[]byte(`[{"width": 111, "coordinates": [[22.2, 33.3], [44.4, 55.5], [66.6, 77.7]]}]`),
			},
			want: []LineString{
				{Width: 111, Coords: []Point{{X: 22.2, Y: 33.3}, {X: 44.4, Y: 55.5}, {X: 66.6, Y: 77.7}}},
			},
			wantErr: false,
		},
		{
			name: "Multiple geometries",
			args: args{
				[]byte(`[
				{"width": 111, "coordinates": [[11.1, 11.1]]},
				{"width": 222, "coordinates": [[22.2, 22.2]]},
				{"width": 333, "coordinates": [[33.3, 33.3]]}
				]`),
			},
			want: []LineString{
				{Width: 111, Coords: []Point{{X: 11.1, Y: 11.1}}},
				{Width: 222, Coords: []Point{{X: 22.2, Y: 22.2}}},
				{Width: 333, Coords: []Point{{X: 33.3, Y: 33.3}}},
			},
			wantErr: false,
		},
		{
			name:    "Zero geometries",
			args:    args{[]byte(`[]`)},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "No data",
			args:    args{[]byte(``)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid data type",
			args:    args{[]byte(`[{"width": "foobar"}]`)},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := geosToLineStrings(test.args.data)
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

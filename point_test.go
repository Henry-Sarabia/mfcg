package mfcg

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPoint_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    Point
		wantErr bool
	}{
		{
			name: "Two float64 elements",
			args: args{
				data: []byte(`[12.3, 45.6]`),
			},
			want:    Point{X: 12.3, Y: 45.6},
			wantErr: false,
		},
		{
			name: "Less than 2 elements",
			args: args{
				data: []byte(`[12.3]`),
			},
			want:    Point{},
			wantErr: true,
		},
		{
			name: "More than 2 elements",
			args: args{
				data: []byte(`[12.3, 45.6, 78.9]`),
			},
			want:    Point{},
			wantErr: true,
		},
		{
			name: "Zero elements",
			args: args{
				data: []byte(`[]`),
			},
			want:    Point{},
			wantErr: true,
		},
		{
			name: "No data",
			args: args{
				data: []byte(``),
			},
			want:    Point{},
			wantErr: true,
		},
		{
			name: "Invalid data type",
			args: args{
				data: []byte(`["12.3","45.6"]`),
			},
			want:    Point{},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got Point
			err := json.Unmarshal(test.args.data, &got)
			if (err != nil) != test.wantErr {
				t.Fatalf("got: <%v>, want error: <%v>", err, test.wantErr)
			}

			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

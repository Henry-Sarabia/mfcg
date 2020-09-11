package mfcg

import "testing"

func TestPoint_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		p       *Point
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "happy path",
			p:       &Point{},
			args:    args{[]byte{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Point.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

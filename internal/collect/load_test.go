package collect

import (
	"testing"
)

func TestParseLoad(t *testing.T) {
	var tests = []struct {
		name    string
		in      string
		want    Load
		wantErr bool
	}{
		{
			name:    "valid",
			in:      "0.65 0.58 0.52 2/1467 29095\n",
			want:    Load{One: 0.65, Five: 0.58, Fifteen: 0.52},
			wantErr: false,
		},
		{
			name:    "too few fields",
			in:      "0.65 0.58",
			wantErr: true,
		},
		{
			name:    "empty",
			in:      "",
			wantErr: true,
		},
		{
			name:    "bad 1 minute",
			in:      "abc 0.58 0.52",
			wantErr: true,
		},
		{
			name:    "bad 5 minute",
			in:      "0.65 abc 0.52",
			wantErr: true,
		},
		{
			name:    "bad 15 minute",
			in:      "0.65 0.58 abc",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLoad(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLoad(%q) error = %v, wantErr %v", tt.in, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("parseLoad(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

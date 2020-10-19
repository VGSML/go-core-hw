package fib

import "testing"

func TestNumber(t *testing.T) {

	tests := []struct {
		name    string
		Num     int
		want    int
		wantErr bool
	}{
		{
			name:    "-1",
			Num:     -1,
			want:    0,
			wantErr: true,
		},
		{
			name:    "3",
			Num:     3,
			want:    2,
			wantErr: false,
		},
		{
			name:    "5",
			Num:     5,
			want:    5,
			wantErr: false,
		},
		{
			name:    "10",
			Num:     10,
			want:    55,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Number(tt.Num)
			if (err != nil) != tt.wantErr {
				t.Errorf("Number() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

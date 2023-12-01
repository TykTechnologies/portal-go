package portal

import (
	"testing"
	"time"
)

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	t0, _ := time.Parse("2006-01-02 15:04", "2020-01-01 00:00")
	tests := []struct {
		name    string
		Time    time.Time
		b       []byte
		wantErr bool
	}{
		{
			name:    "Success",
			Time:    t0,
			b:       []byte("2020-01-01 00:00"),
			wantErr: false,
		},
		{
			name:    "Null-value",
			Time:    t0,
			b:       []byte("null"),
			wantErr: false,
		},
		{
			name:    "Fail",
			Time:    t0,
			b:       []byte("2020-01-01T00:00:00Z00:00"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &CustomTime{
				Time: tt.Time,
			}
			if err := ct.UnmarshalJSON(tt.b); (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want []byte
	}{
		{
			name: "Success",
			time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			want: []byte("2020-01-01 00:00"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &CustomTime{
				Time: tt.time,
			}
			got := ct.MarshalJSON()
			if string(got) != string(tt.want) {
				t.Errorf("CustomTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

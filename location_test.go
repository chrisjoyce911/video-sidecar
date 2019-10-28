package videosidecar

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNewLocation(t *testing.T) {
	type args struct {
		locationName string
	}
	tests := []struct {
		name string
		args args
		want *Location
	}{
		{
			name: "name Moon Rock",
			args: args{
				locationName: "Moon Rock",
			},
			want: &Location{
				LocName:        "Moon Rock",
				LocDisplayName: "moon-rock",
			},
		},
		{
			name: "name Unknown",
			args: args{
				locationName: "",
			},
			want: &Location{
				LocName:        "Unknown",
				LocDisplayName: "unknown",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLocation(tt.args.locationName)
			assert.Equal(t, got, tt.want)
		})
	}
}

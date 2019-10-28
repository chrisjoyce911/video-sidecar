package videosidecar

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNewCountry(t *testing.T) {
	type args struct {
		countryCode string
		countryName string
	}
	tests := []struct {
		name string
		args args
		want *Country
	}{
		{
			name: "name Fantasy code fy",
			args: args{
				countryCode: "fy",
				countryName: "Fantasy",
			},
			want: &Country{
				ID:          "fy",
				CountryName: "Fantasy",
				CountrySlug: "fantasy",
			},
		},
		{
			name: "name Unknown code Unknown",
			args: args{
				countryCode: "",
				countryName: "",
			},
			want: &Country{
				ID:          "zz",
				CountryName: "Unknown",
				CountrySlug: "unknown",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCountry(tt.args.countryCode, tt.args.countryName)
			assert.Equal(t, got, tt.want)
		})
	}
}

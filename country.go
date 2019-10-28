package videosidecar

import (
	"github.com/gosimple/slug"
)

type Country struct {
	ID                 string
	CountrySlug        string
	CountryName        string
	CountryDescription string
	CountryNotes       string
	CountryVideo       *Video
	CountryVideoID     uint
}

// Create a new country
func NewCountry(countryCode string, countryName string) *Country {
	if countryCode == "" {
		countryCode = "zz"
	}

	if countryName == "" {
		countryName = "Unknown"
	}

	countrySlug := slug.MakeLang(countryName, "en")

	result := &Country{
		ID:          countryCode,
		CountryName: countryName,
		CountrySlug: countrySlug,
	}

	return result
}

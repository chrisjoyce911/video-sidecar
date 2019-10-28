package videosidecar

import (
	"strings"

	"github.com/gosimple/slug"
)

// Filmed location
type Location struct {
	LocDisplayName string
	LocLat         float64
	LocLong        float64
	LocCategory    string
	LocType        string
	LocName        string
	LocHouseNr     string
	LocStreet      string
	LocSuburb      string
	LocCity        string
	LocPostcode    string
	LocCounty      string
	LocState       string
	LocCountry     string
	LocCountryCode string
	LocDescription string
	LocNotes       string
	LocFavorite    bool
}

func NewLocation(locationName string) *Location {
	locationName = strings.TrimSpace(locationName)

	if locationName == "" {
		locationName = "Unknown"
	}

	locSlug := slug.Make(locationName)

	result := &Location{
		LocName:        locationName,
		LocDisplayName: locSlug,
	}

	return result
}

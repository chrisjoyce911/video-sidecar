package videosidecar

import (
	"strings"

	"github.com/gosimple/slug"
)

// Labels for video, album and location categorization
type Label struct {
	LabelSlug        string
	LabelName        string
	LabelPriority    int
	LabelFavorite    bool
	LabelDescription string
	LabelNotes       string
	LabelCategories  []*Label
}

func NewLabel(labelName string, labelPriority int) *Label {
	labelName = strings.TrimSpace(labelName)

	if labelName == "" {
		labelName = "Unknown"
	}

	labelSlug := slug.Make(labelName)

	result := &Label{
		LabelName:     labelName,
		LabelSlug:     labelSlug,
		LabelPriority: labelPriority,
	}

	return result
}

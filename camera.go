package videosidecar

import (
	"fmt"
	"strings"

	"github.com/gosimple/slug"
)

// Camera model and make (as extracted from EXIF metadata)
type Camera struct {
	CameraSlug        string
	CameraModel       string
	CameraMake        string
	CameraType        string
	CameraOwner       string
	CameraDescription string
	CameraNotes       string
}

func NewCamera(modelName string, makeName string) *Camera {
	makeName = strings.TrimSpace(makeName)

	if modelName == "" {
		modelName = "Unknown"
	} else if strings.HasPrefix(modelName, makeName) {
		modelName = strings.TrimSpace(modelName[len(makeName):])
	}

	var cameraSlug string

	if makeName != "" {
		cameraSlug = slug.Make(makeName + " " + modelName)
	} else {
		cameraSlug = slug.Make(modelName)
	}

	result := &Camera{
		CameraModel: modelName,
		CameraMake:  makeName,
		CameraSlug:  cameraSlug,
	}

	return result
}

func (m *Camera) String() string {
	if m.CameraMake != "" && m.CameraModel != "" {
		return fmt.Sprintf("%s %s", m.CameraMake, m.CameraModel)
	} else if m.CameraModel != "" {
		return m.CameraModel
	}

	return ""
}

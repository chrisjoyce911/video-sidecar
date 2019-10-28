package videosidecar

// Video labels are weighted by uncertainty (100 - confidence)
type VideoLabel struct {
	VideoID          uint
	LabelID          uint
	LabelUncertainty int
	LabelSource      string
	Video            *Video
	Label            *Label
}

func (VideoLabel) TableName() string {
	return "video_labels"
}

func NewVideoLabel(photoId, labelId uint, uncertainty int, source string) *VideoLabel {
	result := &VideoLabel{
		VideoID:          photoId,
		LabelID:          labelId,
		LabelUncertainty: uncertainty,
		LabelSource:      source,
	}

	return result
}

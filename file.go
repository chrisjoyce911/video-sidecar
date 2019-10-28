package videosidecar

// An video or sidecar file that belongs to a photo
type File struct {
	Video            string
	VideoID          uint
	FileUUID         string
	FileName         string
	FileHash         string
	FileOriginalName string
	FileType         string
	FileMime         string
	FilePrimary      bool
	FileMissing      bool
	FileDuplicate    bool
	FileOrientation  int
	FileNotes        string
}

package videosidecar

import (
	"strings"

	"github.com/gosimple/slug"
	uuid "github.com/satori/go.uuid"
)

// Photo album
type Album struct {
	AlbumUUID        string
	AlbumSlug        string
	AlbumName        string
	AlbumDescription string
	AlbumNotes       string
	AlbumViews       uint
	AlbumVideo       *Video
	AlbumPhotoID     uint
	AlbumFavorite    bool
	Photos           []Video
}

func NewAlbum(albumName string) *Album {
	albumName = strings.TrimSpace(albumName)

	if albumName == "" {
		albumName = "New Album"
	}

	albumSlug := slug.Make(albumName)
	albumUUID := uuid.NewV4().String()

	result := &Album{
		AlbumUUID: albumUUID,
		AlbumSlug: albumSlug,
		AlbumName: albumName,
	}

	return result
}

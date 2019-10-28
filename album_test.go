package videosidecar

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNewAlbum(t *testing.T) {
	type args struct {
		albumName string
	}
	tests := []struct {
		name string
		args args
		want *Album
	}{
		{
			name: "name Christmas 2018",
			args: args{
				albumName: "Christmas 2018",
			},
			want: &Album{
				AlbumName: "Christmas 2018",
				AlbumSlug: "christmas-2018",
			},
		},
		{
			name: "name empty",
			args: args{
				albumName: "",
			},
			want: &Album{
				AlbumName: "New Album",
				AlbumSlug: "new-album",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAlbum(tt.args.albumName)
			assert.Equal(t, got.AlbumName, tt.want.AlbumName)
			assert.Equal(t, got.AlbumSlug, tt.want.AlbumSlug)
		})
	}
}

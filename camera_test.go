package videosidecar

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNewCamera(t *testing.T) {
	type args struct {
		modelName string
		makeName  string
	}
	tests := []struct {
		name string
		args args
		want *Camera
	}{
		{
			name: "model Unknown make Nikon",
			args: args{
				modelName: "",
				makeName:  "Nikon",
			},
			want: &Camera{
				CameraModel: "Unknown",
				CameraMake:  "Nikon",
				CameraSlug:  "nikon-unknown",
			},
		},
		{
			name: "model EOS 6D make Canon",
			args: args{
				modelName: "EOS 6D",
				makeName:  "Canon",
			},
			want: &Camera{
				CameraModel: "EOS 6D",
				CameraMake:  "Canon",
				CameraSlug:  "canon-eos-6d",
			},
		},
		{
			name: "model with prefix make Panasonic",
			args: args{
				modelName: "Panasonic Lumix",
				makeName:  "Panasonic",
			},
			want: &Camera{
				CameraModel: "Lumix",
				CameraMake:  "Panasonic",
				CameraSlug:  "panasonic-lumix",
			},
		},
		{
			name: "model TG-4 make Unknown",
			args: args{
				modelName: "TG-4",
				makeName:  "",
			},
			want: &Camera{
				CameraModel: "TG-4",
				CameraMake:  "",
				CameraSlug:  "tg-4",
			},
		},
		{
			name: "model Unknown make Unknown",
			args: args{
				modelName: "",
				makeName:  "",
			},
			want: &Camera{
				CameraModel: "Unknown",
				CameraMake:  "",
				CameraSlug:  "unknown",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCamera(tt.args.modelName, tt.args.makeName)
			assert.Equal(t, got, tt.want)
		})
	}
}

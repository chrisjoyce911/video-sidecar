package videosidecar

import (
	"reflect"
	"testing"

	"github.com/bmizerany/assert"
)

func TestNewVideoLabel(t *testing.T) {
	type args struct {
		photoId     uint
		labelId     uint
		uncertainty int
		source      string
	}
	tests := []struct {
		name string
		args args
		want *VideoLabel
	}{
		{
			name: "name Unknown",
			args: args{
				photoId:     1,
				labelId:     3,
				uncertainty: 80,
				source:      "source",
			},
			want: &VideoLabel{
				VideoID:          1,
				LabelID:          3,
				LabelUncertainty: 80,
				LabelSource:      "source",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVideoLabel(tt.args.photoId, tt.args.labelId, tt.args.uncertainty, tt.args.source)
			assert.Equal(t, tt.want.VideoID, got.VideoID)
			assert.Equal(t, tt.want.LabelID, got.LabelID)
			assert.Equal(t, tt.want.LabelUncertainty, got.LabelUncertainty)
			assert.Equal(t, tt.want.LabelSource, got.LabelSource)

			if got := NewVideoLabel(tt.args.photoId, tt.args.labelId, tt.args.uncertainty, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVideoLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

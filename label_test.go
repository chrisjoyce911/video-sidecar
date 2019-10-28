package videosidecar

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNewLabel(t *testing.T) {
	type args struct {
		labelName     string
		labelPriority int
	}
	tests := []struct {
		name string
		args args
		want *Label
	}{
		{
			name: "name Unicorn2000 priority 5",
			args: args{
				labelName:     "Unicorn2000",
				labelPriority: 5,
			},
			want: &Label{
				LabelName:     "Unicorn2000",
				LabelPriority: 5,
				LabelSlug:     "unicorn2000",
			},
		},
		{
			name: "name Unknown",
			args: args{
				labelName:     "",
				labelPriority: -6,
			},
			want: &Label{
				LabelName:     "Unknown",
				LabelPriority: -6,
				LabelSlug:     "unknown",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLabel(tt.args.labelName, tt.args.labelPriority)
			assert.Equal(t, got, tt.want)
		})
	}
}

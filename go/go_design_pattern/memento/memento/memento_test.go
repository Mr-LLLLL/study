package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	in := &InputText{}
	snapshot := []*Snapshot{}

	tests := []struct {
		input string
		want  string
	}{
		{
			input: ":list",
			want:  "",
		},
		{
			input: "hello",
			want:  "hello",
		}, {
			input: ":list",
			want:  "hello",
		}, {
			input: "world",
			want:  "",
		}, {
			input: ":undo",
			want:  "",
		}, {
			input: ":list",
			want:  "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			switch tt.input {
			case ":list":
				assert.Equal(t, tt.want, in.GetText())
			case ":undo":
				in.Restore(snapshot[len(snapshot)-1])
				snapshot = snapshot[:len(snapshot)-1]
			default:
				snapshot = append(snapshot, in.Snapshot())
				in.Append(tt.input)
			}
		})
	}
}

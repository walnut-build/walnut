package wfile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSimpleFile(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log(wd)

	assert := assert.New(t)

	data, err := LoadFile("../testdata/wfiles/simple/wfile.yaml")

	assert.Nil(err)
	assert.Nil(data.Include, "Include list should be empty")
	assert.Equal(len(data.Tasks), 2, "Task map should have 2 elements")

	assert.Equal(
		data.Tasks["first"],
		TaskDescription{},
		"Task map should have 2 elements",
	)

	assert.Equal(
		data.Tasks["second"],
		TaskDescription{
			DependsOn: []string{"first"},
		},
		"Task map should have 2 elements",
	)
}

package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code/parsers"
)

func TestBuild_ProducesExpectedDiff(t *testing.T) {
	data1 := parsers.ParsedData{
		"common": parsers.ParsedData{
			"setting1": "Value 1",
			"setting2": 200,
			"setting3": true,
			"setting6": parsers.ParsedData{
				"key": "value",
				"deep": parsers.ParsedData{
					"wow": "",
				},
			},
		},
	}
	data2 := parsers.ParsedData{
		"common": parsers.ParsedData{
			"follow":   false,
			"setting1": "Value 1",
			"setting3": nil,
			"setting4": "blah blah",
			"setting5": parsers.ParsedData{
				"key5": "value5",
			},
			"setting6": parsers.ParsedData{
				"key": "value",
				"ops": "vops",
				"deep": parsers.ParsedData{
					"wow": "so much",
				},
			},
		},
	}

	got := Build(data1, data2)

	assert.Len(t, got, 1)
	assert.Equal(t, "common", got[0].Key)
	assert.Equal(t, Nested, got[0].Status)

	commonChildren := got[0].Children
	assert.Len(t, commonChildren, 7)
	assert.Equal(t, "follow", commonChildren[0].Key)
	assert.Equal(t, Added, commonChildren[0].Status)
	assert.Equal(t, "setting1", commonChildren[1].Key)
	assert.Equal(t, Unchanged, commonChildren[1].Status)
	assert.Equal(t, "setting2", commonChildren[2].Key)
	assert.Equal(t, Deleted, commonChildren[2].Status)
	assert.Equal(t, "setting3", commonChildren[3].Key)
	assert.Equal(t, Changed, commonChildren[3].Status)
	assert.Equal(t, "setting4", commonChildren[4].Key)
	assert.Equal(t, Added, commonChildren[4].Status)
	assert.Equal(t, "setting5", commonChildren[5].Key)
	assert.Equal(t, Added, commonChildren[5].Status)
	assert.Equal(t, "setting6", commonChildren[6].Key)
	assert.Equal(t, Nested, commonChildren[6].Status)

	nestedChildren := commonChildren[6].Children
	assert.Len(t, nestedChildren, 3)
	assert.Equal(t, "deep", nestedChildren[0].Key)
	assert.Equal(t, Nested, nestedChildren[0].Status)
	assert.Equal(t, "key", nestedChildren[1].Key)
	assert.Equal(t, Unchanged, nestedChildren[1].Status)
	assert.Equal(t, "ops", nestedChildren[2].Key)
	assert.Equal(t, Added, nestedChildren[2].Status)

	deepChildren := nestedChildren[0].Children
	assert.Len(t, deepChildren, 1)
	assert.Equal(t, "wow", deepChildren[0].Key)
	assert.Equal(t, Changed, deepChildren[0].Status)
}

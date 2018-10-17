package processor_test

import (
	"errors"
	"github.com/rady-io/annotation-processor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessor_Scan(t *testing.T) {
	i := 0
	err := processor.NewProcessor(`
//
@Page(name) {name}
@Body json
@File(avatar)
@Deprecated
@NeverBeHere
`).Scan(func(ann, key, value string) (err error) {
		i++
		count := i
		switch count {
		case 1:
			assert.Equal(t, "@Page", ann)
			assert.Equal(t, "name", key)
			assert.Equal(t, "{name}", value)
		case 2:
			assert.Equal(t, "@Body", ann)
			assert.Equal(t, "", key)
			assert.Equal(t, "json", value)
		case 3:
			assert.Equal(t, "@File", ann)
			assert.Equal(t, "avatar", key)
			assert.Equal(t, "", value)
		case 4:
			assert.Equal(t, "@Deprecated", ann)
			assert.Equal(t, "", key)
			assert.Equal(t, "", value)
			err = errors.New("Unsupported annotation: " + ann)
		}
		return
	})

	assert.Error(t, err)
	assert.Equal(t, "Unsupported annotation: @Deprecated", err.Error())
	assert.Equal(t, 4, i)
}

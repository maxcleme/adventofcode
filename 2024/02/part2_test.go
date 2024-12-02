package _02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	f := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	reports, err := parse(f)
	assert.NoError(t, err)
	assert.Equal(t, reports, []Report{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	})
	assert.Equal(t, reports[0].safeWithDampener(), true)
	assert.Equal(t, reports[1].safeWithDampener(), false)
	assert.Equal(t, reports[2].safeWithDampener(), false)
	assert.Equal(t, reports[3].safeWithDampener(), true)
	assert.Equal(t, reports[4].safeWithDampener(), true)
	assert.Equal(t, reports[5].safeWithDampener(), true)
}

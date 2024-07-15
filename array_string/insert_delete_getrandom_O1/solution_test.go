package insert_delete_getrandom_O1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedSet(t *testing.T) {
	set := Constructor()

	insert := set.Insert(1)
	assert.Equal(t, true, insert)

	remove := set.Remove(2)
	assert.Equal(t, false, remove)

	insert = set.Insert(2)
	assert.Equal(t, true, insert)

	random := set.GetRandom()
	_ = random

	remove = set.Remove(1)
	assert.Equal(t, true, remove)

	insert = set.Insert(2)
	assert.Equal(t, false, insert)

	random = set.GetRandom()
	fmt.Println(random)
}

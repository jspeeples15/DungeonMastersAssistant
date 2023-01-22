package characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LoadParty(t *testing.T) {
	p, err := LoadParty("../../configs", "TestParty")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(p.Players))
	assert.Equal(t, "The example friends", p.Name)
}

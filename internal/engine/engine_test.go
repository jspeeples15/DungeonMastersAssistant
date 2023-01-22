package engine

import (
	"DungeonMastersAssistant/internal/characters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEngine_Run(t *testing.T) {
	e := Engine{
		WorkingDirectory: "../../configs",
		Party:            characters.Party{},
	}
	e.Run()
}

func TestEngine_HandleInput(t *testing.T) {
	e := Engine{
		WorkingDirectory: "../../configs",
		Party:            characters.Party{},
	}

	err := e.HandleInput([]string{"load", "TestParty"})

	assert.Equal(t, 2, len(e.Party.Players))
	assert.Nil(t, err)
}

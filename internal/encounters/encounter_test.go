package encounters

import (
	"DungeonMastersAssistant/internal/characters"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_Encounter_GetInitiativeOrder(t *testing.T) {
	player1 := characters.Player{
		Dex: 12,
	}

	player2 := characters.Player{
		Dex: 14,
	}

	player3 := characters.Player{
		Dex: 8,
	}

	player4 := characters.Player{
		Dex: 10,
	}

	player5 := characters.Player{
		Dex: 10,
	}

	encounter := Encounter{participants: []characters.Character{player1, player2, player3, player4, player5}}

	result := encounter.GetInitiativeOrder()

	assert.Equal(t, 5, len(result))
}

func TestEncounter_Sort(t *testing.T) {
	player1 := InitiativeHaver{characters.Player{
		Dex: 12,
	}, 19,
	}

	player2 := InitiativeHaver{characters.Player{
		Dex: 12,
	}, 11,
	}

	player3 := InitiativeHaver{characters.Player{
		Dex: 12,
	}, 12,
	}

	player4 := InitiativeHaver{characters.Player{
		Dex: 12,
	}, 16,
	}

	player5 := InitiativeHaver{characters.Player{
		Dex: 12,
	}, 14,
	}

	ih := InitiativeHavers{player1, player2, player3, player4, player5}

	sort.Sort(ih)

	assert.Equal(t, player1, ih[0])
	assert.Equal(t, player4, ih[1])
	assert.Equal(t, player5, ih[2])
	assert.Equal(t, player3, ih[3])
	assert.Equal(t, player2, ih[4])
}

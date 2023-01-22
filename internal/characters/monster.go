package characters

import (
	"DungeonMastersAssistant/internal/util"
	"fmt"
)

type Monster struct {
	Name          string   `json:"name"`
	HP            int      `json:"hp"`
	Str           int      `json:"str"`
	Dex           int      `json:"dex"`
	Con           int      `json:"con"`
	Wis           int      `json:"wis"`
	Int           int      `json:"int"`
	Cha           int      `json:"cha"`
	MovementSpeed int      `json:"movementSpeed"`
	Actions       []string `json:"actions"`
}

func (m Monster) GetInitiative() (initiativeScore int) {
	roll := util.Roll(20)
	modifier := util.CalculateModifier(m.Dex)

	return roll + modifier
}

func (m Monster) Print() {
	fmt.Sprintf("Monster: %s [ %v ]", m.Name, m.HP)
}

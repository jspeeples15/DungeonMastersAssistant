package characters

import (
	"DungeonMastersAssistant/internal/listener"
	"DungeonMastersAssistant/internal/util"
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
	Name          string `json:"name"`
	HP            int    `json:"hp"`
	Str           int    `json:"str"`
	Dex           int    `json:"dex"`
	Con           int    `json:"con"`
	Wis           int    `json:"wis"`
	Int           int    `json:"int"`
	Cha           int    `json:"cha"`
	movementSpeed int    `json:"movementSpeed"`
}

func CreatePlayerViaPrompts() (p *Player, err error) {
	//must define player
	p = &Player{}

	//name
	resp, err := listener.Prompt("New Player Name? ")
	if err != nil {
		fmt.Printf("Error with New Name Response: %s", err.Error())
		return p, err
	}
	var sb strings.Builder
	for _, v := range resp {
		sb.WriteString(v)
	}
	name := sb.String()

	//hp
	resp, err = listener.Prompt("New Player Max HP? ")
	if err != nil {
		fmt.Printf("Error with New HP Response: %s", err.Error())
		return p, err
	}
	if len(resp) > 1 {
		fmt.Println("Error with New HP Input: Invalid Number of Inputs. Must just be 1 integer")
		return p, err
	}

	maxHP, convErr := strconv.Atoi(resp[0])
	if convErr != nil {
		fmt.Printf("Error with New HP Input: Invalid Number. %s \n", resp[0])
		return p, err
	}

	//todo: get other stats

	p.SetName(name)
	p.SetMaxHP(maxHP)
	return p, nil
}

func (p *Player) SetName(newName string) {
	//todo: debug statement only
	fmt.Printf("Setting player name to: %s \n", newName)

	p.Name = newName
}

func (p *Player) SetMaxHP(maxHP int) {
	//todo: debug statement only
	//fmt.Printf("Setting player max hp to: %v \n", maxHP)

	p.HP = maxHP
}
func (p *Player) GetInitiative() int {
	roll := util.Roll(20)
	modifier := util.CalculateModifier(p.Dex)

	return roll + modifier
}

func (p *Player) Print() {
	fmt.Printf("%s [%v] \n", p.Name, p.HP)
}

func (p *Player) Damage(dmg int) {
	p.HP = p.HP - dmg
}

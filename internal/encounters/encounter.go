package encounters

import (
	"DungeonMastersAssistant/internal/characters"
	"sort"
)

type Encounter struct {
	participants []characters.Character
}

func NewEncounter() Encounter {
	return Encounter{}
}
func (e *Encounter) LoadParticipants() {
	players := e.LoadPlayers()
	monsters := e.LoadMonsters()
	e.participants = append(players, monsters...)
}

func (e *Encounter) LoadPlayers() (cList []characters.Character) {
	players := []characters.Player{}
	for _, player := range players {
		cList = append(cList, player)
	}
	return cList
}

func (e *Encounter) LoadMonsters() (cList []characters.Character) {
	monsters := []characters.Monster{}
	for _, monster := range monsters {
		cList = append(cList, monster)
	}
	return cList
}
func (e *Encounter) GetInitiativeOrder() []characters.Character {
	initHavers := InitiativeHavers{}
	for _, participant := range e.participants {
		init := participant.GetInitiative()

		initHavers = append(initHavers, InitiativeHaver{
			participant: participant,
			initiative:  init,
		})
	}
	sort.Sort(initHavers)

	return initHavers.GetParticipantsFromInit()
}

type InitiativeHaver struct {
	participant characters.Character
	initiative  int
}

type InitiativeHavers []InitiativeHaver

func (ih InitiativeHavers) GetParticipantsFromInit() []characters.Character {
	results := []characters.Character{}
	for _, v := range ih {
		results = append(results, v.participant)
	}

	return results
}

func (ih InitiativeHavers) Len() int {
	return len(ih)
}

func (ih InitiativeHavers) Less(i, j int) bool {
	return ih[i].initiative > ih[j].initiative
}

func (ih InitiativeHavers) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

package characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Party struct {
	Name    string    `json:"name"`
	Players []*Player `json:"players"`
}

func LoadParty(configDirectory, name string) (players Party, err error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("working directory!!!!!!!")
	fmt.Println(path)

	fullPath := fmt.Sprintf("%s/parties/%s.json", configDirectory, name)
	fmt.Println(fmt.Sprintf("Full loading path: %s", fullPath))
	jsonFile, err := os.Open(fullPath)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening party file: %s", err.Error()))
		return players, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &players)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error reading party file: %s", err.Error()))
		return players, err
	}
	return players, nil
}

func (p *Party) SetPlayers(players []*Player) {
	p.Players = players
}

func (p *Party) AddPlayer(player *Player) {
	if p.Players == nil {
		p.Players = []*Player{player}
	} else {
		p.Players = append(p.Players, player)
	}
}

func (p *Party) SaveParty(workingDirectory string) (err error) {
	file, err := json.Marshal(p)
	if err != nil {
		return err
	}

	writeLoc := fmt.Sprintf("%s/parties/%s.json", workingDirectory, p.Name)
	err = ioutil.WriteFile(writeLoc, file, 0644)

	return nil
}

func (p *Party) Print() {
	fmt.Println("Current Party:")
	fmt.Printf("Name: %s", p.Name)
	fmt.Println("")
	fmt.Println("Characters:")
	for _, v := range p.Players {
		v.Print()
	}
	fmt.Println("")
}

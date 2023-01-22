package engine

import (
	"DungeonMastersAssistant/internal/characters"
	"DungeonMastersAssistant/internal/listener"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Engine struct {
	WorkingDirectory string
	Party            characters.Party
}

type Input []string

func (e *Engine) Run() {
	fmt.Println("Welcome to DND-CLI!")
	fmt.Println("To Load Existing Party: Press 1 or type 'Load {Party_Name}'")
	fmt.Println("To Create New Party: Press 2 or type 'New [optional]{Party_Name}'")

	//load party
	input, err := listener.Listen()
	if err != nil {
		fmt.Sprintf("ERROR WITH INPUT: %s. I am sorry I failed you, user", err.Error())
		os.Exit(1)
	}

	fmt.Println("Great input! It says:")
	for _, v := range input {
		fmt.Print(v)
		fmt.Print(",")
	}
	err = e.HandleInput(input)
	if err != nil {
		fmt.Sprintf("ERROR WITH ACTION ROUTING: %s. I am sorry I failed you, user", err.Error())
		os.Exit(1)
	}

	//todo: listen for further commands
	if len(e.Party.Players) == 0 {
		fmt.Println("NO PLAYERS LOADED YET?!! THATS NOT A-SUPPOSED TO HAPPEN! BYE BYE")
		os.Exit(1)
	}

	e.Party.Print()
	fmt.Println("Great Job!")
	fmt.Println("CONSTRUCTION ZONE XX==XX==XX==XX LEAVE NOW")
}

func (e *Engine) HandleInput(input []string) (err error) {
	action_input := input[0]
	fmt.Printf("Action Input: %s \n", action_input)
	action := DecipherActionInput(action_input)

	//todo: if undefined, retry
	switch action {
	case CREATE_PARTY:
		return e.CreateNewParty(input)
	case LOAD_PARTY:
		return e.LoadParty(input)
	default:
		fmt.Print("Invalid Input: ")
		for _, v := range input {
			fmt.Printf("%s ", v)
		}
		return errors.New("Invalid Input")
	}
}

func (e *Engine) CreateNewParty(input []string) (err error) {
	if len(input) < 2 {
		errMsg := "Invalid Input Length For New Party: Must be 'new {party name}'"
		fmt.Println(errMsg)
		return errors.New(errMsg)
	}

	newParty := characters.Party{}
	newParty.Name = input[1]

	keepGoing := true
	for keepGoing {
		p, err := characters.CreatePlayerViaPrompts()
		if err != nil {
			fmt.Printf("Error Creating Player, %s \n", err.Error())
			return err
		}
		newParty.AddPlayer(p)

		resp, err := listener.Prompt("Success! Add Another Party Member? (y/n): ")
		isValidInputResp := e.isValidContinueResp(resp)
		if err != nil {
			fmt.Printf("Error Reading Input, Closing Party Editor. %s \n", err.Error())
			return err
		} else if !isValidInputResp {
			for !isValidInputResp {
				fmt.Println("INVALID INPUT! OPTIONS ['y', 'yes', 'n', 'no']")
				resp, err := listener.Prompt("Add Another Party Member? (y/n): ")
				if err != nil {
					fmt.Printf("Error Reading Input, Closing Party Editor. %s \n", err.Error())
					return err
				}
				isValidInputResp = e.isValidContinueResp(resp)
			}
		}

		//todo: cleaner way to check this- a lot of non-DRY substrings
		if resp[0][:1] == "n" {
			keepGoing = false
		}
	}

	e.Party = newParty
	err = newParty.SaveParty(e.WorkingDirectory)
	if err != nil {
		fmt.Printf("Error Saving Party: %s \n", err.Error())
		return err
	}
	return nil
}

func (e *Engine) isValidContinueResp(input []string) bool {
	if len(input) != 1 {
		//too many words
		return false
	}
	//just check the first letter buddy
	propResp := strings.ToLower(input[0])[:1]

	return propResp == "y" || propResp == "n"
}

func (e *Engine) LoadParty(input []string) (err error) {
	if len(input) < 2 {
		errMsg := "Invalid Input Length For Existing Party: Must be 'load {party name}'"
		fmt.Println(errMsg)
		return errors.New(errMsg)
	}
	fmt.Println("we are about to load a party")
	fmt.Println("working directory: " + e.WorkingDirectory)
	fmt.Println("party name: " + input[1])
	e.Party, err = characters.LoadParty(e.WorkingDirectory, input[1])
	if err != nil {
		fmt.Sprintf("Error Loading Party %s. %s", input[1], err.Error())
	}
	return err
}

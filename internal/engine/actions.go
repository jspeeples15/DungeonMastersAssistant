package engine

import "strings"

type Action string

const (
	LOAD_PARTY   Action = "load_party"
	CREATE_PARTY Action = "create_party"
	UNDEFINED    Action = "undefined"
)

func DecipherActionInput(input string) Action {
	input = strings.ToLower(input)
	switch input {
	case "load":
		return LOAD_PARTY
	case "create":
		return CREATE_PARTY
	case "new":
		return CREATE_PARTY
	default:
		return UNDEFINED
	}
}

func (a Action) GetString() string {
	return string(a)
}

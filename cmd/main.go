package main

import (
	"DungeonMastersAssistant/internal/engine"
)

func main() {
	e := engine.Engine{
		WorkingDirectory: "configs",
	}
	e.Run()
}

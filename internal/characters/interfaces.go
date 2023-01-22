package characters

type Character interface {
	GetInitiative() (initiativeScore int)
	Print()
}

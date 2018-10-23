package main

type Characters struct {
	dndcs []DnDCharacter
	// TODO: add support for more games
}

func (collection Characters) addDnd(char DnDCharacter) {
	collection.dndcs = append(collection.dndcs, char)
}

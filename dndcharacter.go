package main

const (
	ATHLETICS = iota
	ACROBATICS
	SLEIGHTOFHAND
	STEALTH
	ARCANA
	HISTORY
	INVESTIGATION
	NATURE
	RELIGION
	ANIMALHANDLING
	INSIGHT
	MEDICINE
	PERCEPTION
	SURVIVAL
	DECEPTION
	INTIMIDATION
	PERFORMANCE
	PERSUASION
	SKILLS
)

const (
	STR = iota
	DEX
	CON
	INT
	WIS
	CHA
	STATS
)

type DnDCharacter struct {
	name string
	//todo: race and class
	stats       [STATS]int
	inventory   map[string]int
	money       map[string]int
	hp          int
	ac          int
	bonus       int
	skills      [SKILLS]int
	pp          int
	saving      [STATS]int
	attBonusStr int
	attBonusDex int
	initiative  int
	speed       int
}

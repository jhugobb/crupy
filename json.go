package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func saveDndCharacter(c DnDCharacter) {

	marshalledJSON, _ := json.Marshal(c)
	err := ioutil.WriteFile("characters/dnd/"+c.name+".json", marshalledJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func buildAllDnd() string {
	b := strings.Builder{}
	files, _ := ioutil.ReadDir("characters/dnd/")
	for i, f := range files {
		jsonChar, err := ioutil.ReadFile(f.Name())
		if err != nil {
			log.Fatal(err)
		}
		var character DnDCharacter
		err = json.Unmarshal(jsonChar, &character)
		if err != nil {
			log.Fatal(err)
		}
		b.WriteString("Character #" + strconv.Itoa(i+1) + ": " + character.name + " is a [race] [class]\n")
	}
	return b.String()
}

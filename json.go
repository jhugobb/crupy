package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func saveDndCharacter(c *DnDCharacter) {

	marshalledJSON, _ := json.Marshal(c)
	err := ioutil.WriteFile("characters/dnd/"+c.name+".json", marshalledJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func buildAllDnd() string {
	b := strings.Builder{}
	pwd := "characters/dnd/"
	files, _ := ioutil.ReadDir(pwd)
	for i, f := range files {
		isJSON := filepath.Ext(pwd+f.Name()) == ".json"
		if isJSON {
			log.Print(f.Name())
			jsonChar, err := ioutil.ReadFile(pwd + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			log.Print(string(jsonChar))
			var character DnDCharacter
			err = json.Unmarshal(jsonChar, &character)
			if err != nil {
				log.Fatal(err)
			}
			b.WriteString("Character #" + strconv.Itoa(i+1) + ": " + character.name + " is a [race] [class]\n")
		}
	}
	return b.String()
}

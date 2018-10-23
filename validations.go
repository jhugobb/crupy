package main

import (
	"errors"
	"log"
	"regexp"
)

func validateCreate(msg []string) error {
	var err error
	switch msg[0] {
	case "dnd":
		break
	default:
		err = errors.New("The rpg you want to create a character is not supported or not existant. Sorry :/")
		return err
	}
	isok, err := regexp.MatchString("^[a-zA-Z]+$", msg[1])
	if err != nil {
		log.Fatal(err)
	}
	if !isok {
		err = errors.New("Oops! Looks like your name is not valid! Try and use only letters")
		return err
	}
	return nil
}

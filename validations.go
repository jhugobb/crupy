package main

import (
	"errors"
	"regexp"
)

func validateCreate(msg []string) error {
	var err error
	// TODO: verify that the number of arguments is correct
	switch msg[0] {
	case "dnd":
	default:
		err = errors.New("The rpg you want to create a character is not supported or not existant. Sorry :/")
		return err
	}
	isok, _ := regexp.MatchString("^[a-zA-Z]+$", msg[1])
	if !isok {
		err = errors.New("Oops! Looks like your name is not valid! Try and use only letters")
		return err
	}
	return nil
}

func validateShowAll(msg []string) error {
	var err error
	if len(msg) != 1 {
		err = errors.New("Your command syntax is not correct! Show all command usage: /showall [game]")
		return err
	}
	switch msg[0] {
	case "dnd":
	default:
		err = errors.New("The rpg you want to show all character is not supported or not existant. Sorry :/")
		return err
	}
	return nil
}

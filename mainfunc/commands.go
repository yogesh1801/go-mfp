// MFP      - Miulti-Function Printers and scanners toolkit
// mainfunc - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// List of all commands

package mainfunc

// Command contains a command's description
type Command struct {
	Name string              // Command name
	Main func(args []string) // Command main function
}

// CommandByName finds command by name.
// If command is not found, it returns nil.
func CommandByName(name string) *Command {
	for i := range Commands {
		if name == Commands[i].Name {
			return &Commands[i]
		}
	}

	return nil
}

// Commands contains descriptions of all implemented commands
var Commands = []Command{
	{Name: "cups", Main: MainCups},
}

package models

type Command struct {
	CommandType string
	Args        []string
}

type SETCommand struct {
	Key   string
	Value string
}

type GETCommand struct {
	Key   string
	Value string
}

var Commands = []Command{}
var SETCommands = []SETCommand{}

var SETHashCommand = make(map[string]string)

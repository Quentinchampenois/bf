package main

import (
	"fmt"
)

const filename = ".brainflood"

type Brainflood struct {
	Global BrainfloodGlobal `toml:"global"`
	Path   string
}

type BrainfloodGlobal struct {
	Name        string
	Description string
	Language    string
	Tags        []string
	Author      string
}

func (b Brainflood) GetInfo() {
	fmt.Println(b.Global.Name)
	fmt.Println(b.Global.Description)
	fmt.Println(b.Global.Language)
	fmt.Println(b.Global.Tags)
	fmt.Println(b.Global.Author)
}

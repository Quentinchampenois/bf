package main

import (
	"fmt"
)

const filename = ".brainflood"

type Bf struct {
	Global BfGlobal `toml:"global"`
	Path   string
}

type BfGlobal struct {
	Name        string
	Description string
	Language    string
	Tags        []string
	Author      string
}

func (b Bf) GetInfo() {
	fmt.Println(b.Global.Name)
	fmt.Println(b.Global.Description)
	fmt.Println(b.Global.Language)
	fmt.Println(b.Global.Tags)
	fmt.Println(b.Global.Author)
}

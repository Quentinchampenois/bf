package main

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
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
	License     string
	Author      string
}

func (b Brainflood) GetInfo() {
	fmt.Println(b.Global.Name)
	fmt.Println(b.Global.Description)
	fmt.Println(b.Global.Language)
	fmt.Println(b.Global.Tags)
	fmt.Println(b.Global.License)
	fmt.Println(b.Global.Author)
}

func Projects() {
	entries, err := os.ReadDir(projectsDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		filepath := fmt.Sprintf("%s/%s/%s", projectsDir, e.Name(), filename)
		file, err := os.ReadFile(filepath)
		if err != nil {
			continue
		}

		if len(file) == 0 {
			fmt.Println(fmt.Sprintf("%s : File is empty", filepath))
			continue
		}

		brainflood := Brainflood{
			Path: filepath,
		}
		err = toml.Unmarshal(file, &brainflood)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		brainflood.GetInfo()
	}
}

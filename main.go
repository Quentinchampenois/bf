package main

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

var homeDir = os.Getenv("HOME")
var projectsDir = fmt.Sprintf("%s/Documents/local", homeDir)

func main() {
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

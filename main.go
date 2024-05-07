package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"strings"
)

var homeDir = os.Getenv("HOME")
var projectsDir = fmt.Sprintf("%s/Documents/local", homeDir)

func main() {
	fmt.Println(homeDir)
	fmt.Println(projectsDir)

	app := &cli.App{
		Name:  "new",
		Usage: "create a new brainflood project",
		Action: func(*cli.Context) error {
			currentPath, _ := os.Getwd()

			bf := Brainflood{
				Path: currentPath,
			}
			prompt := promptui.Prompt{
				Label:    "Project name",
				Validate: Validate(3),
			}

			result, _ := prompt.Run()
			bf.Global.Name = result

			prompt = promptui.Prompt{
				Label:    "Description",
				Validate: Validate(8),
			}

			result, _ = prompt.Run()
			bf.Global.Description = result

			promptSelect := promptui.Select{
				Label: "Programming language",
				Items: []string{"Go", "Python", "JavaScript", "Ruby", "Rust", "Java", "Kotlin", "Swift", "C", "C++"},
				Size:  10,
			}

			_, result, _ = promptSelect.Run()

			bf.Global.Language = result

			AppendToFile(".registry", currentPath)
			b, err := toml.Marshal(bf)
			if err != nil {
				log.Fatal(err)
			}

			AppendToFile(fmt.Sprintf("%s/%s", currentPath, ".brainflood"), string(b))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	//Projects()
}

func AppendToFile(filename, data string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// No need to append path if it already exists
	if strings.Contains(string(content), data) {
		return
	}

	if _, err := f.Write([]byte(fmt.Sprintf("%s\n", data))); err != nil {
		fmt.Println(err.Error())
	}
}

func Validate(minChars int) func(s string) error {
	return func(s string) error {
		if len(s) < minChars {
			return fmt.Errorf("Name must be at least %d characters", minChars)
		}

		if s[0] < 65 || s[0] > 90 {
			return fmt.Errorf("Name must start with a capital letter")
		}
		return nil
	}
}

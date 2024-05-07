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

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "create a new brainflood project",
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
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all brainflood projects",
				Action: func(*cli.Context) error {
					currentPath, _ := os.Getwd()

					bytes, err := ReadFile(fmt.Sprintf("%s/%s", currentPath, ".registry"))
					if err != nil {
						log.Fatal(err)
					}

					paths := strings.Split(string(bytes), "\n")
					projects := make([]Brainflood, 0, len(paths)-1)
					for _, p := range paths {
						if p == "" {
							continue
						}
						content, err := ReadFile(fmt.Sprintf("%s/%s", p, ".brainflood"))
						if err != nil {
							fmt.Println("Errors reading")
							fmt.Println(err.Error())
							continue
						}

						brainflood := Brainflood{
							Path: p,
						}

						err = toml.Unmarshal(content, &brainflood)
						if err != nil {
							fmt.Println("Errors")
							fmt.Println(err.Error())
							continue
						}

						projects = append(projects, brainflood)
					}

					if len(projects) < 1 {
						fmt.Println("No projects found")
						return nil
					}
					fmt.Println(fmt.Sprintf("Found %d projects", len(projects)))

					names := make([]string, 0, len(projects))
					for _, p := range projects {
						names = append(names, p.Global.Name)
					}
					promptSelect := promptui.Select{
						Label: "Projects",
						Items: names,
						Size:  len(names),
					}

					_, result, _ := promptSelect.Run()
					for _, p := range projects {
						if p.Global.Name == result {
							fmt.Println(p.Global.Description)
							fmt.Println(p.Path)
						}
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func ReadFile(filepath string) ([]byte, error) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)

	return bytes, err
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

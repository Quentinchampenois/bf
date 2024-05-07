# Bf

Simple tool to centralize projects on the filesystem. I have plenty of projects on my computer and I wanted to have a simple way to access them. I also wanted to have a way to quickly create new projects.

## How it works

Thanks to the `bf` command, you can create a new project, list all your projects, and open them in your favorite editor.

When a new project is created, it add an entry in file `.bf-registry`. This file is located in the root directory of the project. It contains the name of the project and the path to the project.
Thanks to this entry, we can find quickly the project on the filesystem and the purpose.

## Usage

1. Define `BF_REGISTRY` environment variable to the path of the registry file. 
2. Clone project
3. Compile project : `go build -o bf`

## CLI

Available commands

* `bf new` : Create a new project
* `bf list` : List all projects

## Todo 

* [ ] Add a way to remove a project
* [ ] Add a way to edit a project
* [ ] Do not add entry in `.registry` if the project is already created
* [ ] Add default date to current on create
* [ ] Add a way to list all projects with a specific tag
* [ ] Add a way to list all projects by tags
* [ ] Add a way to list all projects by language
* [ ] Add unit testing
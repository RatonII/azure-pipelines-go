package main

type PipelinesConfig struct {
	Organization string `yaml:"organization"`
	Pipelines Pipelines
}
type Pipelines []Pipeline

type Pipeline struct {
	Project string		`yaml:"project"`
	Name string			`yaml:"name"`
	Description string  `yaml:"description"`
	Folder string		`yaml:"folder,omitempty"`
	Repository string 	`yaml:"repository,omitempty"`
	Branch string	  	`yaml:"branch,omitempty"`
	YamlPath	string	`yaml:"yamlpath"`
}

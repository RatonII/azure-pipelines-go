package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

func CreateFolderPipeline(folderName string, project string) string {
	if proc, err := Start( "az", "pipelines", "folder", "create",
		"--project", project,
		"--path", folderName,
		"--description", "Folder for WARP project",); err == nil {
		proc.Wait()
	}
	return folderName
}

func CreatePipeline(project string, name string,
					description string, folder string,
					repository string, branch string,
					yamlpath string) {
	if proc, err := Start( "az", "pipelines", "create",
		"--project", project,
		"--name", name,
		"--description", description,
		"--folder", CreateFolderPipeline(folder, project),
		"--repository",  repository,
		"--branch", branch,
		"--repository-type", "tfsgit",
		"--yml-path", yamlpath); err == nil {
		proc.Wait()
	}
}
func  DevOpsLogin(org string) {
	if proc, err := Start( "az",
		"devops",
		"login",
		"--org", org); err == nil {
		proc.Wait()
	}
}

func (c *Pipelines) getConf(PipelinesFile *string) *Pipelines {

	yamlFile, err := ioutil.ReadFile(*PipelinesFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

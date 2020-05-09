package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
		pipelinesFile := flag.String("file","","Add a config file yaml with all the pipelines contains")
		flag.Parse()
		if *pipelinesFile != "" {
			var pipelines Pipelines
			fmt.Println(pipelines.getConf(pipelinesFile))
			for _, pipeline := range *pipelines.getConf(pipelinesFile) {
				DevOpsLogin(pipeline.Organization)
				CreatePipeline(	pipeline.Project,
								pipeline.Name,
								pipeline.Description,
								pipeline.Folder,
								pipeline.Repository,
								pipeline.Branch,
								pipeline.YamlPath)
			}
		}else {
			log.Fatalln("Please specify a config file for the pipelines with the argument --file")
		}



//	organizationUrl := "https://dev.azure.com/mariuss2007"  // todo: replace value with your organization url
//	personalAccessToken := "5trqrjcf2c2phqil6jmgm2foy775yw77lp4ak4c7pojmlhl4j42a"  // todo: replace value with your PAT
//
//	// Create a connection to your organization
//	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
//	project := "test"
//	//folder := "TestFolder"
//	pipelineName := "NewTestPipeline"
//	repository := pipelines.RepositoryTypeValues.AzureReposGit
//	repositoryName := "test"
//	guid := "1"
//	ctx := context.Background()
//	client := pipelines.NewClient(ctx,connection)
//	conftype := pipelines.ConfigurationTypeValues.Yaml
//
//	// Get first page of the list of team projects for your organization
//	responseValue, err := client.ListPipelines(ctx, pipelines.ListPipelinesArgs{
//		Project: &project,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	createValue, err := client.CreatePipeline(ctx, pipelines.CreatePipelineArgs{
//		Project: &project,
//		InputParameters: &pipelines.CreatePipelineParameters{
//			Configuration: &pipelines.CreatePipelineConfigurationParameters{
//				Type: &conftype,
//				Path: "azure-pipelines.yaml",
//				Repository: &pipelines.Repository{
//					Type: &repository,
//					RepositoryId: &repositoryName,
//					Guid: &guid },
//				RepositoryId: &repositoryName,
//			},
//		Name:          &pipelineName,
//		},
//	})
//	if err != nil {
//		panic(err)
//	}
//	log.Println(*createValue)
//	// Print response from pipeline creation
//		//log.Printf("Name = %v",*createValue)
//
//	// Log the page of team project names
//		for _, teamProjectReference := range (*responseValue).Value {
//			log.Printf("Name = %v",*teamProjectReference.Id)
//		}
}
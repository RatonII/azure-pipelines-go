package main

import (
	"flag"
	"fmt"
	"log"

	"sync"
)

func main() {
		var wg sync.WaitGroup
		pipelinesFile := flag.String("file","","Add a config file yaml with all the pipelines contains")
		flag.Parse()
		if *pipelinesFile != "" {
			var p PipelinesConfig
			fmt.Println(p.getConf(pipelinesFile).Organization)
			pipelines := p.getConf(pipelinesFile).Pipelines
			org := p.getConf(pipelinesFile).Organization
			DevOpsLogin(org)
			for _, pipeline := range pipelines {
				wg.Add(1)
				go CreatePipeline(	pipeline.Project,
					pipeline.Name,
					pipeline.Description,
					pipeline.Folder,
					pipeline.Repository,
					pipeline.Branch,
					pipeline.YamlPath,&wg)
			}
			wg.Wait()
		}else {
			log.Fatalln("Please specify a config file for the pipelines with the argument --file")
		}
}
package main

import (
	"flag"
	"log"
	"runtime"

	"sync"
)

func main() {
		var wg sync.WaitGroup
		runtime.GOMAXPROCS(4)
		pipelinesFile := flag.String("file","","Add a config file yaml with all the pipelines contains")
		flag.Parse()
		if *pipelinesFile != "" {
			var p PipelinesConfig
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
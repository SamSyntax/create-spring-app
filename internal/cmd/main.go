package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SamSyntax/create-spring-app/internal/core"
	"github.com/SamSyntax/create-spring-app/internal/fetcher"
	"github.com/SamSyntax/create-spring-app/internal/logger"
	"github.com/SamSyntax/create-spring-app/internal/misc"
	"github.com/charmbracelet/log"
)

func main() {
	log.SetReportTimestamp(false)
	log.SetLevel(log.DebugLevel)
	conf, err := core.CreateProjectConfig()
	if err != nil {
		log.Error(fmt.Printf("Couldn't create a project config: %e\n", err))
		os.Exit(1)
	}
	dir, err := misc.GetWorkingFolder()
	if err != nil {
		log.Error(fmt.Printf("Couldn't create a project config: %e\n", err))
	}

	err = core.RunForm(conf)
	if err != nil {
		log.Error(fmt.Printf("Couldn't run: %e\n", err))
		os.Exit(1)
	}
	apiUrl := core.BuildUrl(*conf)
	split := strings.Split(dir, "/")
	logger.LogInfo(fmt.Sprintf("Initializing project at %s...\n", split[len(split)-1]))
	err = fetcher.DownloadAndExtract(apiUrl, conf.ArtifactId)
	if err != nil {
		log.Error(fmt.Printf("Failed: %v\n", err))
		os.Exit(1)
	}
	logger.LogSuccess(fmt.Sprintf("Success! Project created in %s\n", split[len(split)-1]))

}

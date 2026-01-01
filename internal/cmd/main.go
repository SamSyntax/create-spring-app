package main

import (
	"fmt"
	"github.com/SamSyntax/create-spring-app/internal/core"
	"github.com/SamSyntax/create-spring-app/internal/logger"
	"github.com/SamSyntax/create-spring-app/internal/misc"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"net/http"
	"os"
	"os/exec"
	"strings"
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
	}
	apiUrl := core.BuildUrl(*conf)
	split := strings.Split(dir, "/")
	logger.LogInfo(fmt.Sprintf("Initializing project at %s...\n", split[len(split)-1]))
	err = downloadAndExtract(apiUrl, conf.ArtifactId)
	if err != nil {
		log.Error(fmt.Printf("Failed: %v\n", err))
		os.Exit(1)
	}
	logger.LogSuccess(fmt.Sprintf("Success! Project created in %s\n", split[len(split)-1]))

}

func downloadAndExtract(targetUrl string, dir string) error {
	errChan := make(chan error, 1)

	go func() {
		os.MkdirAll(dir, 0755)
		res, err := http.Get(targetUrl)
		if err != nil {
			errChan <- err
			return
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			errChan <- fmt.Errorf("API error: %s", res.Status)
			return
		}

		cmd := exec.Command("tar", "-xzf", "-", "-C", ".")
		cmd.Stdin = res.Body
		err = cmd.Run()
		errChan <- err
	}()

	p := tea.NewProgram(misc.SpinnerModel{
		Spinner: spinner.New(spinner.WithSpinner(spinner.Pulse)),
		ErrChan: errChan,
	})

	_, err := p.Run()

	return err

}

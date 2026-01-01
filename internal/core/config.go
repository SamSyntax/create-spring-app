package core

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/SamSyntax/create-spring-app/internal/fetcher"
)

type ProjectConfig struct {
	Build             string
	Language          string
	JavaVersion       string
	ArtifactId        string
	SpringBootVersion fetcher.Val
	Dependencies      []string
	Meta              *fetcher.InitMetadata
	PackageName       string
	GroupName         string
}

func CreateProjectConfig() (*ProjectConfig, error) {
	meta, err := fetcher.FetchDependencies()
	if err != nil {
		return nil, err
	}

	return &ProjectConfig{
		Build:       "Maven",
		Language:    "Java",
		JavaVersion: "21",
		SpringBootVersion: fetcher.Val{
			ID:   meta.BootVersion.Default,
			Name: meta.BootVersion.Default,
		},
		Dependencies: nil,
		Meta:         meta,
	}, nil
}

func BuildUrl(conf ProjectConfig) string {
	baseUrl := "https://start.spring.io/starter.tgz"
	bootVersion := strings.ReplaceAll(conf.SpringBootVersion.ID, ".BUILD", "")
	fmt.Println(bootVersion, conf.SpringBootVersion.ID)
	params := url.Values{}
	params.Add("type", conf.Build)
	params.Add("artifactId", conf.ArtifactId)
	params.Add("baseDir", conf.ArtifactId)
	params.Add("javaVersion", conf.JavaVersion)
	params.Add("bootVersion", conf.SpringBootVersion.ID)
	params.Add("groupId", conf.GroupName)
	params.Add("packageName", conf.PackageName)
	params.Add("dependencies", strings.Join(conf.Dependencies, ","))
	fmt.Printf("%s?%s\n", baseUrl, params.Encode())
	return fmt.Sprintf("%s?%s", baseUrl, params.Encode())
}

package core

import (
	"fmt"
	"strings"

	"github.com/SamSyntax/create-spring-app/internal/fetcher"
	"github.com/SamSyntax/create-spring-app/internal/misc"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

func (pc *ProjectConfig) CreateDepsForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title("Dependencies").Options(
				pc.Meta.DepJsonToHuh(pc.SpringBootVersion.ID)...,
			).Value(&pc.Dependencies).Filterable(true).Validate(func(selected []string) error {
				for _, id := range selected {
					dep := pc.Meta.FindDependencyByID(id)
					if dep != nil && !fetcher.CheckCompatibility(pc.SpringBootVersion.ID, dep.VersionRange) {
						return fmt.Errorf("'%s' is not compatible with Spring Boot %s\n", dep.Name, pc.SpringBootVersion.Name)
					}
				}
				return nil
			}),
		),
	)
}

func RunForm(pc *ProjectConfig) error {
	theme := GetCustomTheme()
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Project Name").Value(&pc.ArtifactId).Description("Enter Project Name").Validate(misc.ValidateNoSpaces)),
		huh.NewGroup(
			huh.NewInput().Title("Group Name").Value(&pc.GroupName).Placeholder("com.example").Description("Enter Group Name").Validate(misc.ValidateNoSpaces)),
	).WithTheme(theme).Run()
	if pc.GroupName == "" {
		pc.GroupName = "com.example"
	}
	var placeholder string
	if pc.ArtifactId == "" {
		dir, _ := misc.GetWorkingFolder()
		splits := strings.Split(dir, "/")
		name := splits[len(splits)-1]
		placeholder = fmt.Sprintf("%s.%s", pc.GroupName, name)
	} else {
		placeholder = fmt.Sprintf("%s.%s", pc.GroupName, pc.ArtifactId)
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Package Name").Value(&pc.PackageName).Placeholder(placeholder).Description("Enter package name").Validate(misc.ValidateNoSpaces)),

		huh.NewGroup(
			huh.NewSelect[fetcher.Val]().Title("SpringBoot version").Options(
				pc.Meta.BootVersionJsonToHuh()...,
			).Value(&pc.SpringBootVersion),
		),

		huh.NewGroup(
			huh.NewSelect[string]().Title("Build Tool").Options(
				huh.NewOption("Maven", "maven-project"),
				huh.NewOption("Gradle - Groovy", "gradle-project"),
				huh.NewOption("Gradle - Kotlin", "gradle-project-kotlin"),
			).Value(&pc.Build)),

		huh.NewGroup(
			huh.NewSelect[string]().Title("Java Version").Options(
				pc.Meta.JavaVersionJsonToHuh()...,
			).Value(&pc.JavaVersion)),
	).WithTheme(theme).Run()
	if err != nil {
		return err
	}

	depForm := pc.CreateDepsForm()
	if err := depForm.Run(); err != nil {
		log.Fatal(err)
	}
	return nil
}

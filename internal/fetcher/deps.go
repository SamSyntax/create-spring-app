package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	dimmedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	errorStyle  = lipgloss.NewStyle().Foreground(lipgloss.CompleteColor{
		TrueColor: "#a01000",
	})
)

type InitMetadata struct {
	Dependencies struct {
		Values []struct {
			Values []struct {
				ID           string `json:"id"`
				Name         string `json:"name"`
				VersionRange string `json:"versionRange"`
			} `json:"values"`
		} `json:"values"`
	} `json:"dependencies"`

	JavaVersion struct {
		Default string `json:"default"`
		Values  []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"javaVersion"`

	BootVersion struct {
		Default string `json:"default"`
		Values  []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"bootVersion"`
}

func (md *InitMetadata) DepJsonToHuh(selectedBootVersion string) []huh.Option[string] {
	var options []huh.Option[string]
	if len(md.Dependencies.Values) == 0 {
		return nil
	}
	for _, group := range md.Dependencies.Values {
		for _, dep := range group.Values {
			label := dep.Name
			fmt.Println(selectedBootVersion, dep.VersionRange)
			isComp := CheckCompatibility(selectedBootVersion, dep.VersionRange)
			if !isComp {
				label = errorStyle.Render("✗ ") + dimmedStyle.Render(dep.Name) +
					dimmedStyle.Render(fmt.Sprintf(" [Needs: %s]", dep.VersionRange))
			} else {
				label = dep.Name
			}
			options = append(options, huh.NewOption(label, dep.ID))
		}
	}

	return options
}

type Val struct {
	ID           string
	Name         string
	VersionRange string
}

func (md *InitMetadata) BootVersionJsonToHuh() []huh.Option[Val] {
	var options []huh.Option[Val]
	if len(md.BootVersion.Values) == 0 {
		return nil
	}
	for _, ver := range md.BootVersion.Values {
		options = append(options, huh.NewOption(ver.Name, Val{
			Name: ver.Name,
			ID:   ver.ID,
		}))
	}
	return options
}
func (md *InitMetadata) JavaVersionJsonToHuh() []huh.Option[string] {
	var options []huh.Option[string]
	if len(md.JavaVersion.Values) == 0 {
		return nil
	}
	for _, ver := range md.JavaVersion.Values {
		options = append(options, huh.NewOption(ver.Name, ver.ID))
	}

	return options
}

func (md InitMetadata) FindDependencyByID(id string) *Val {
	for _, group := range md.Dependencies.Values {
		for _, dep := range group.Values {
			if id == dep.ID {
				return &Val{
					ID:           dep.ID,
					Name:         dep.Name,
					VersionRange: dep.VersionRange,
				}
			}
		}
	}
	return nil
}

func FetchDependencies() (*InitMetadata, error) {
	req, _ := http.NewRequest("GET", "https://start.spring.io/metadata/client", nil)
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var meta InitMetadata

	if err := json.NewDecoder(res.Body).Decode(&meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

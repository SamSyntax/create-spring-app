package core

import (
	"fmt"

	"github.com/SamSyntax/create-spring-app/internal/fetcher"
)

func (pc *ProjectConfig) ValidateDependencies(selected []string) error {
	for _, id := range selected {
		dep := pc.Meta.FindDependencyByID(id)
		if dep != nil && !fetcher.CheckCompatibility(pc.SpringBootVersion.ID, dep.VersionRange) {
			return fmt.Errorf("'%s' requires Spring Boot %s", dep.Name, dep.VersionRange)
		}
	}
	return nil
}

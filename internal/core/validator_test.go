package core

import (
	"encoding/json"
	"testing"

	"github.com/SamSyntax/create-spring-app/internal/fetcher"
)

func TestProjectConfig_ValidateDependencies(t *testing.T) {
	metaJson := `
	{
		"dependencies": {
			"values": [
				{
					"values": [
						{
							"id": "web",
							"name": "Spring Web",
							"versionRange": ""
						},
						{
							"id": "security",
							"name": "Spring Security",
							"versionRange": "[2.0.0.RELEASE,3.0.0.RELEASE)"
						},
						{
							"id": "data-jpa",
							"name": "Spring Data JPA",
							"versionRange": ">=2.5.0.RELEASE"
						}
					]
				}
			]
		}
	}
	`
	var meta fetcher.InitMetadata
	if err := json.Unmarshal([]byte(metaJson), &meta); err != nil {
		t.Fatalf("Failed to unmarshal test metadata: %v", err)
	}

	tests := []struct {
		name          string
		bootVersion   string
		dependencies  []string
		wantErr       bool
	}{
		{
			name:         "Compatible dependencies",
			bootVersion:  "2.6.0",
			dependencies: []string{"web", "security", "data-jpa"},
			wantErr:      false,
		},
		{
			name:         "Incompatible security",
			bootVersion:  "3.1.0",
			dependencies: []string{"security"},
			wantErr:      true,
		},
		{
			name:         "Incompatible data-jpa (too low)",
			bootVersion:  "2.4.0",
			dependencies: []string{"data-jpa"},
			wantErr:      true,
		},
		{
			name:         "Unknown dependency",
			bootVersion:  "2.6.0",
			dependencies: []string{"unknown"},
			wantErr:      false, // Logic ignores unknown deps currently: "dep != nil" check
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProjectConfig{
				SpringBootVersion: fetcher.Val{ID: tt.bootVersion},
				Meta:              &meta,
			}
			err := pc.ValidateDependencies(tt.dependencies)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

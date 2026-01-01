package core

import (
	"net/url"
	"strings"
	"testing"

	"github.com/SamSyntax/create-spring-app/internal/fetcher"
)

func TestBuildUrl(t *testing.T) {
	tests := []struct {
		name string
		conf ProjectConfig
		want map[string]string
	}{
		{
			name: "Standard Maven Project",
			conf: ProjectConfig{
				Build:       "maven-project",
				ArtifactId:  "demo-service",
				JavaVersion: "21",
				GroupName:   "com.example",
				PackageName: "com.example.demo",
				SpringBootVersion: fetcher.Val{
					ID: "3.4.1.RELEASE",
				},
				Dependencies: []string{"web", "data-jpa"},
			},
			want: map[string]string{
				"type":         "maven-project",
				"artifactId":   "demo-service",
				"baseDir":      "demo-service",
				"javaVersion":  "21",
				"bootVersion":  "3.4.1.RELEASE",
				"groupId":      "com.example",
				"packageName":  "com.example.demo",
				"dependencies": "web,data-jpa",
			},
		},
		{
			name: "Gradle Project No Dependencies",
			conf: ProjectConfig{
				Build:       "gradle-project",
				ArtifactId:  "minimal",
				JavaVersion: "17",
				SpringBootVersion: fetcher.Val{
					ID: "3.2.0",
				},
				Dependencies: []string{},
			},
			want: map[string]string{
				"type":         "gradle-project",
				"artifactId":   "minimal",
				"dependencies": "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildUrl(tt.conf)

			u, err := url.Parse(got)
			if err != nil {
				t.Fatalf("Failed to parse generated URL: %v", err)
			}

			if !strings.HasPrefix(got, "https://start.spring.io/starter.tgz") {
				t.Errorf("BuildUrl() base path = %v, want prefix https://start.spring.io/starter.tgz", u.Path)
			}

			q := u.Query()
			for key, wantVal := range tt.want {
				gotVal := q.Get(key)
				if gotVal != wantVal {
					t.Errorf("Query param %s = %v, want %v", key, gotVal, wantVal)
				}
			}
		})
	}
}

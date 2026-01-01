package main

import (
	"net/url"
	"testing"

	"github.com/SamSyntax/create-spring-app/internal/core"
	"github.com/SamSyntax/create-spring-app/internal/fetcher"
)

func Test_buildUrl(t *testing.T) {
	conf := core.ProjectConfig{
		Build:       "maven-project",
		ArtifactId:  "demo",
		JavaVersion: "17",
		SpringBootVersion: fetcher.Val{
			ID: "3.2.0",
		},
		GroupName:    "com.example",
		PackageName:  "com.example.demo",
		Dependencies: []string{"web", "actuator"},
	}

	got := core.BuildUrl(conf)
	u, err := url.Parse(got)
	if err != nil {
		t.Fatalf("buildUrl returned invalid URL: %v", err)
	}

	if u.Scheme != "https" || u.Host != "start.spring.io" || u.Path != "/starter.tgz" {
		t.Errorf("Unexpected base URL: %s", u.String())
	}

	q := u.Query()
	tests := []struct {
		key  string
		want string
	}{
		{"type", "maven-project"},
		{"artifactId", "demo"},
		{"baseDir", "demo"},
		{"javaVersion", "17"},
		{"bootVersion", "3.2.0"},
		{"groupId", "com.example"},
		{"packageName", "com.example.demo"},
		{"dependencies", "web,actuator"},
	}

	for _, tt := range tests {
		if gotVal := q.Get(tt.key); gotVal != tt.want {
			t.Errorf("Query param %s = %s, want %s", tt.key, gotVal, tt.want)
		}
	}
}

func BuildUrl(conf core.ProjectConfig) any {
	panic("unimplemented")
}

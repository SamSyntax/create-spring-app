package fetcher

import (
	"strings"

	"github.com/Masterminds/semver/v3"
)

func normalizeVersion(v string) string {
	v = strings.ReplaceAll(v, ".RELEASE", "")
	v = strings.ReplaceAll(v, ".M", "-M")
	v = strings.ReplaceAll(v, ".RC", "-RC")
	v = strings.ReplaceAll(v, ".BUILD-SNAPSHOT", "-SNAPSHOT")
	v = strings.ReplaceAll(v, ".SNAPSHOT", "-SNAPSHOT")

	return v
}

func CheckCompatibility(bootVersion, springRange string) bool {
	if springRange == "" {
		return true
	}

	v, err := semver.NewVersion(normalizeVersion(bootVersion))
	if err != nil {
		return false
	}

	springRange = strings.ReplaceAll(springRange, " ", "")
	parts := strings.Split(springRange, ",")

	if len(parts) != 2 {
		norm := normalizeVersion(springRange)
		c, err := semver.NewConstraint(norm)
		if err == nil {
			return c.Check(v)
		}
		c, err = semver.NewConstraint(">= " + norm + "-0")
		if err != nil {
			return false
		}
		return c.Check(v)
	}

	left := parts[0]
	right := parts[1]

	lowVer := normalizeVersion(strings.Trim(left, "[]()"))
	highVer := normalizeVersion(strings.Trim(right, "[]()"))

	var sb strings.Builder

	if strings.HasPrefix(left, "[") {
		sb.WriteString(">= ")
	} else {
		sb.WriteString("> ")
	}
	sb.WriteString(lowVer + "-0")
	sb.WriteString(", ")

	if strings.HasSuffix(right, "]") {
		sb.WriteString("<= ")
	} else {
		sb.WriteString("< ")
	}
	sb.WriteString(highVer)

	c, err := semver.NewConstraint(sb.String())
	if err != nil {
		return false
	}

	return c.Check(v)
}

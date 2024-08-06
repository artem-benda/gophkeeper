package main

import "strings"

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

// VersionString - получить строковое представление версии приложения
func VersionString() string {
	var sb strings.Builder
	sb.WriteString("Build version: ")
	if buildVersion != "" {
		sb.WriteString(buildVersion)
	} else {
		sb.WriteString("N/A")
	}

	sb.WriteString(", Build date: ")
	if buildDate != "" {
		sb.WriteString(buildDate)
	} else {
		sb.WriteString("N/A")
	}

	sb.WriteString(", Build commit: ")
	if buildCommit != "" {
		sb.WriteString(buildCommit)
	} else {
		sb.WriteString("N/A")
	}

	return sb.String()
}

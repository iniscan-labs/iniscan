package main

import (
	"embed"
	"fmt"
	"log/slog"
	"regexp"

	"github.com/iniscan-labs/iniscan/cmd"
	"github.com/iniscan-labs/iniscan/pkg"
)

//go:embed CHANGELOG.md
var changeLogFS embed.FS

// @title iniscan API
// @version 0.0.1
// @description iniscan api document
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath
func main() {
	getLatestVersion()

	cmd.Execute()
}

func getLatestVersion() {
	changeLog, err := changeLogFS.ReadFile("CHANGELOG.md")
	if err != nil {
		slog.Warn("read CHANGELOG.md failed", slog.Any("err", err))
		return
	}
	// Regular expression to match version numbers
	versionRegex := regexp.MustCompile(`## v(\d+\.\d+\.\d+)`)

	// Find all matches in the changelog
	matches := versionRegex.FindAllSubmatch(changeLog, -1)

	if len(matches) == 0 {
		fmt.Println("No version numbers found")
		return
	}

	// The first match corresponds to the latest version since the changelog lists them in descending order
	pkg.Version = string(matches[0][1])
}

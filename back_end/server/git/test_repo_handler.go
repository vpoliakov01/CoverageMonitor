package git

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/vpoliakov01/CoverageMonitor/back_end/server/parser"
	"github.com/vpoliakov01/CoverageMonitor/back_end/server/utils"
)

// CloneRepo clones repo into the specified location
func CloneRepo(org, repo, path string) error {
	url := fmt.Sprintf("https://github.com/%v/%v.git", org, repo)
	args := []string{
		"clone",
		url,
		path,
	}
	return exec.Command("git", args...).Run()
}

// TestRepo tests the repo with go test and returns the coverage
func TestRepo(path string) (*parser.ProjectCoverage, error) {
	args := []string{
		"-c",
		fmt.Sprintf("cd %v && go test -coverprofile coverage.out", path),
	}
	err := exec.Command("bash", args...).Run()
	if err != nil {
		return nil, err
	}

	coveragePath := filepath.Join(path, "coverage.out")
	coverage, err := ioutil.ReadFile(coveragePath)
	if err != nil {
		return nil, err
	}
	pc := parser.ParseCover(coverage)
	return &pc, nil
}

// TestRepoHandler returns the test coverage statistics
func TestRepoHandler(c *gin.Context) {
	org := c.Param("github_org")
	repo := c.Param("github_repo")

	if org == "" || repo == "" {
		utils.Abort(c, http.StatusBadRequest, "invalid request")
		return
	}

	dir := filepath.Join("repos", org, repo)
	err := CloneRepo(org, repo, dir)
	if err != nil {
		utils.Abort(c, http.StatusInternalServerError, fmt.Sprintf("cloning repo failed: %v", err))
		return
	}
	defer os.RemoveAll(dir)

	coverage, err := TestRepo(dir)
	if err != nil {
		utils.Abort(c, http.StatusInternalServerError, fmt.Sprintf("testing repo failed: %v", err))
		return
	}

	utils.Reply(c, http.StatusOK, "OK", coverage)
}

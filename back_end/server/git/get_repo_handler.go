package git

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vpoliakov01/CoverageMonitor/back_end/server/utils"
)

// GetRepoFiles returns the repo files
func GetRepoFiles(org, repo string) ([]FileInfo, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.github.com/repos/%v/%v/contents", org, repo),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := []FileInfo{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetFile returns the file from the repo
func GetFile(org, repo, path string) (*FileInfo, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", org, repo, path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := FileInfo{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRepoMeta returns the number of repo watchers and repo language
func GetRepoMeta(org, repo string) (*RepoMeta, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.github.com/repos/%v/%v", org, repo),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := RepoMeta{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRepo the gets the repo
func GetRepo(org, repo string) (*Repo, error) {
	files, err := GetRepoFiles(org, repo)
	if err != nil {
		return nil, err
	}

	ch := make(chan error)
	// Populate files' content
	for i := range files {
		go func(i int) { // Run calls in parallel
			fileInfo, err := GetFile(org, repo, files[i].Path)
			if err != nil {
				ch <- err
			}
			files[i] = *fileInfo
			ch <- nil
		}(i)
	}
	for range files {
		err = <-ch
		if err != nil {
			return nil, err
		}
	}

	meta, err := GetRepoMeta(org, repo)
	if err != nil {
		return nil, err
	}

	return &Repo{
		Name:  repo,
		Org:   org,
		Meta:  meta,
		Files: files,
	}, nil
}

// GetRepoHandler returns repo files info
func GetRepoHandler(c *gin.Context) {
	org := c.Param("github_org")
	repo := c.Param("github_repo")

	if org == "" || repo == "" {
		utils.Abort(c, http.StatusBadRequest, "invalid request")
		return
	}

	r, err := GetRepo(org, repo)
	if err != nil {
		utils.Abort(c, http.StatusInternalServerError, "getting repo failed")
		return
	}
	utils.Reply(c, http.StatusOK, "OK", r)
}

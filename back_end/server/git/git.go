package git

// FileInfo contains needed info about the file
type FileInfo struct {
	Path    string `json:"path"`
	Content []byte `json:"content"` // base64-encoded
}

// RepoMeta contains repo meta info
type RepoMeta struct {
	Watchers int    `json:"watchers"`
	Language string `json:"language"`
}

// Repo represents a git repo
type Repo struct {
	Name  string     `json:"name"`
	Org   string     `json:"org"`
	Meta  *RepoMeta  `json:"meta"`
	Files []FileInfo `json:"files"` // base64-encoded
}

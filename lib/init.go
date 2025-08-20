package lib

import (
	"os"
	"path/filepath"
)

func NewInitConfig(path string) *InitConfig {
	return &InitConfig{
		Path: path,
	}
}

func (i InitConfig) Init() error {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(filepath.Join(i.Path, dir), 0755); err != nil {
			return err		
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(filepath.Join(i.Path, ".git/HEAD"), headFileContents, 0644); err != nil {
		return err
	}

	return nil
}

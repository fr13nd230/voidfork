package lib

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func NewCatFileConfig(path string) *CatFileConfig {
	return &CatFileConfig{
		InitPath: path,
	}
}

func (cf CatFileConfig) CatFile(hash string, flag string) error {
	if len(hash) < SHA1_LENGTH {
		return errors.New("SHA-1 hash has invalid length unable to continue.")
	}

	objectDir := ".git/objects"
	fileDir := hash[:2]
	fileName := hash[2:]

	_, err := os.ReadDir(filepath.Join(cf.InitPath, objectDir, fileDir))
	if err != nil {
		return err
	}

	f, err := os.Open(filepath.Join(cf.InitPath, objectDir, fileDir, fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := zlib.NewReader(f)
	if err != nil {
		return err
	}
	defer r.Close()

	// var buf bytes.Buffer
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	
	nullIdx := bytes.IndexByte(data, 0)
	switch flag {
		case "-p": // Stand for pretty print
			blob := data[nullIdx+1:]
			fmt.Println(strings.TrimSpace(string(blob)))
		// TODO: In future add other flags parsing.
		default:
			return errors.New("Invalid argument " + flag)
	}	
	return nil
}

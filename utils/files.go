package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name     string
	Path     string
	Size     int64
	FullPath string
	Hash     string
}

func ListFiles(dir string) ([]*File, error) {
	var files []*File
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			var f File
			f.Size = info.Size()

			fields := strings.Split(strings.TrimPrefix(strings.TrimPrefix(path, dir), "/"), "/")

			f.Path = strings.Join(fields[:len(fields)-1], "/")
			f.Name = fields[len(fields)-1]
			f.FullPath = path

			fo, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fo.Close()

			h := sha256.New()
			if _, err := io.Copy(h, fo); err != nil {
				return err
			}

			f.Hash = fmt.Sprintf("%x", h.Sum(nil))

			files = append(files, &f)

			return nil
		})

	return files, err
}

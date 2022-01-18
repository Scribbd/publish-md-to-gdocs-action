package api

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

type Local struct {
	Src string
	Rec bool
}

func (x Local) GetPaths() ([]string, error) {
	var paths []string
	// Get path strings for later usage
	if x.Rec { // When recursing use the advanced WalkDir()
		filepath.WalkDir(x.Src, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				paths = append(paths, path)
			}
			return nil
		})
	} else { // Else use simple ReadDir
		fileInfo, err := ioutil.ReadDir(x.Src)
		if err != nil {
			return paths, err
		}
		for _, file := range fileInfo {
			if !file.IsDir() {
				paths = append(paths, x.Src+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return paths, nil
}

func GetDoc(path string) (ast.Node, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	buf = markdown.NormalizeNewlines(buf)
	parsed := markdown.Parse(buf, parser.New())

	return parsed, nil
}

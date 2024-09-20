package shell

import (
	"archive/zip"
	"os"
)

type (
	FileSystem interface {
		GetCurPathString() string
	}
	fileSystem struct {
		zw      *zip.Writer
		curPath []string
	}
)

func newFS(apath string) (fileSystem, error) {
	a, err := os.OpenFile(apath, os.O_CREATE, 0770)
	if err != nil {
		return fileSystem{}, err
	}

	fs := fileSystem{
		zw:      zip.NewWriter(a),
		curPath: []string{},
	}

	return fs, nil
}

func (fs fileSystem) GetCurPathString() string {
	spath := "/"
	for _, d := range fs.curPath {
		spath += d + "/"
	}

	return spath
}

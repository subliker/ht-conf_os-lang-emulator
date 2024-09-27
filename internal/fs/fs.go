package fs

import (
	"archive/zip"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/artdarek/go-unzip"
)

type (
	// go:generate go run https://github.com/vektra/mockery/v2@v2.46.0 --name = FileSystem
	FileSystem interface {
		CurPath() string
		WriteFile(name string, rw bool, data string) error
		OpenFile(name string) (*os.File, error)
		ChangeDirectory(name string) error
		List(write func(string), more bool) error
		WriteZip() error
		Clear()
	}
	fileSystem struct {
		apath   string
		fd      string
		curPath string
		init    bool
	}
)

var (
	ErrFSNotInit = errors.New("filesystem wasn't initialized")
)

func NewFS(apath string) (FileSystem, error) {
	td := os.TempDir()
	if td == "nil" {
		panic("error finding temp directory")
	}

	fd, err := os.MkdirTemp(td, "")
	if err != nil {
		panic("error creating temp directory: " + err.Error())
	}

	uz := unzip.New(apath, fd)
	uz.Extract()

	fs := fileSystem{
		apath:   apath,
		fd:      fd,
		curPath: string(filepath.Separator),
		init:    true,
	}

	return &fs, nil
}

func (fs *fileSystem) path(path string) string {
	return filepath.Join(fs.fd, path)
}

func (fs *fileSystem) CurPath() string {
	return fs.curPath
}

func (fs *fileSystem) WriteFile(name string, rw bool, data string) error {
	if !fs.init {
		return ErrFSNotInit
	}

	flags := os.O_CREATE
	if rw {
		flags |= os.O_TRUNC
	} else {
		flags |= os.O_APPEND
	}

	_, err := os.Stat(filepath.Dir(fs.path(filepath.Join(fs.curPath, name))))
	if os.IsNotExist(err) {
		return errors.New(name + ": No such file or directory")
	} else if err != nil {
		return err
	}

	f, err := os.OpenFile(fs.path(filepath.Join(fs.curPath, name)), flags, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		return err
	}
	return nil
}

func (fs *fileSystem) OpenFile(name string) (*os.File, error) {
	if !fs.init {
		return nil, ErrFSNotInit
	}

	f, err := os.Open(fs.path(filepath.Join(fs.curPath, name)))
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (fs *fileSystem) ChangeDirectory(path string) error {
	if !fs.init {
		return ErrFSNotInit
	}

	ap, err := filepath.Abs(filepath.Join(fs.path(fs.curPath), path))
	if err != nil {
		return err
	}

	i, err := os.Stat(ap)
	if os.IsNotExist(err) {
		return errors.New("directory " + filepath.Join(fs.curPath, path) + " wasn't found")
	} else if i != nil && !i.IsDir() {
		return errors.New(filepath.Join(fs.curPath, path) + " is not directory")
	} else if err != nil {
		return err
	}
	fs.curPath = filepath.Join(fs.curPath, path)
	return nil
}

func (fs *fileSystem) List(write func(string), more bool) error {
	if !fs.init {
		return ErrFSNotInit
	}

	files, err := os.ReadDir(fs.path(fs.curPath))
	if err != nil {
		return err
	}

	for _, f := range files {
		s := ""
		n := f.Name()
		if f.IsDir() {
			n = "\033[34m" + n + "\033[0m"
		}
		if more {
			i, err := os.Stat(filepath.Join(fs.path(fs.curPath), f.Name()))
			if err != nil {
				s = n
			} else {
				if f.IsDir() {
					s += "d "
				} else {
					s += "- "
				}
				s += strconv.FormatInt(i.Size(), 10) + " " + i.ModTime().Format(time.RFC822) + " " + n + "\n"
			}
		} else {
			s += n + "\n"
		}
		write(s)
	}
	return nil
}

func (fs *fileSystem) WriteZip() error {
	if !fs.init {
		return ErrFSNotInit
	}

	f, err := os.OpenFile(fs.apath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	w := zip.NewWriter(f)
	defer w.Close()

	w.AddFS(os.DirFS(fs.fd))
	return nil
}

func (fs *fileSystem) Clear() {
	if err := os.RemoveAll(fs.fd); err != nil {
		panic("error clearing temp data")
	}
}

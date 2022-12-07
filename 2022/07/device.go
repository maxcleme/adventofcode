package _07

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
)

type FS struct {
	path    string
	parent  *FS
	size    int
	content map[string]*FS
}

func NewFS(path string) *FS {
	return &FS{
		path:    path,
		content: map[string]*FS{},
	}
}

func (fs FS) String() string {
	s := fs.path
	if fs.size != 0 {
		s += fmt.Sprintln(" ", fs.size)
	}
	for _, v := range fs.content {
		s += v.String()
	}
	return s
}

func (fs *FS) Walk(fn func(path string, dir bool, size int)) {
	fn(fs.path, len(fs.content) != 0, fs.size)
	for _, v := range fs.content {
		v.Walk(fn)
	}
}

func (fs *FS) Root() *FS {
	if fs.parent == nil {
		return fs
	}
	return fs.parent.Root()
}

func (fs *FS) Dirs() []*FS {
	var dirs []*FS
	for _, fs := range fs.content {
		if len(fs.content) != 0 {
			dirs = append(dirs, fs)
		}
	}
	return dirs
}

func (fs *FS) Parse(s string) *FS {
	cmdRgx := regexp.MustCompile(`^\$\s+(\w+)\s?(.*)$`)
	fileRgx := regexp.MustCompile(`^(\d+) (.*)$`)
	dirRgx := regexp.MustCompile(`^dir (.*)$`)
	switch {
	case cmdRgx.MatchString(s):
		group := cmdRgx.FindAllStringSubmatch(s, -1)
		cmd := group[0][1]
		args := group[0][2:]
		fs = ParseCommand(fs, cmd, args...)
	case fileRgx.MatchString(s):
		group := fileRgx.FindAllStringSubmatch(s, -1)
		size := group[0][1]
		file := group[0][2]
		fs.content[file] = NewFS(filepath.Join(fs.path, file))
		fs.content[file].size, _ = strconv.Atoi(size)
		fs.size += fs.content[file].size
		for p := fs.parent; p != nil; p = p.parent {
			p.size += fs.content[file].size
		}
	case dirRgx.MatchString(s):
		dir := dirRgx.FindAllStringSubmatch(s, -1)[0][1]
		if _, ok := fs.content[dir]; !ok {
			fs.content[dir] = NewFS(filepath.Join(fs.path, dir))
			fs.content[dir].parent = fs
		}
	}
	return fs
}

func ParseCommand(fs *FS, cmd string, args ...string) *FS {
	switch cmd {
	case "ls":
		return fs
	case "cd":
		return ParseCD(fs, args[0])
	}
	return fs
}

func ParseCD(fs *FS, path string) *FS {
	switch path {
	case "/":
		return fs.Root()
	case "..":
		return fs.parent
	default:
		return fs.content[path]
	}
}

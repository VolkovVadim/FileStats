package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type FileInfo struct {
	Path string
	Size uint64
}

func (fi FileInfo) Update() error {

	return nil
}

type DirInfo struct {
	Path  string
	Size  uint64
	Dirs  []DirInfo
	Files []FileInfo
}

func NewDirInfo(path string) *DirInfo {
	return &DirInfo{
		Path: path,
	}
}

func (di DirInfo) Update() error {
	fmt.Printf("Update : %s\n", di.Path)

	return nil
}

type Stats struct {
	RootPath    string
	RootDirInfo DirInfo
}

func NewStats(path string) *Stats {
	return &Stats{
		RootPath: path,
	}
}

func (s Stats) Collect() error {
	s.RootDirInfo = *NewDirInfo(s.RootPath)

	err := s.RootDirInfo.Update()
	if err != nil {
		return errors.Wrap(err, "unable to update root dir")
	}

	return nil
}

func (s Stats) Show() {
	fmt.Printf("Total size : %d", s.RootDirInfo.Size)
}

func main() {
	fmt.Println("FileStats v0.1")

	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}

	RootPath := filepath.Dir(execPath)
	fmt.Printf("Root path : %s\n", RootPath)

	CommonStats := NewStats(RootPath)
	err = CommonStats.Collect()
	if err != nil {
		fmt.Printf("Error : unablr to collect stats : %v\n", err)
		return
	}

	fmt.Println("Done")
}

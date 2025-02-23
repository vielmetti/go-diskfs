package main

import (
	"fmt"
	"os"

	"github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
)

func main() {
	filename := "test_file.img"
	os.Remove(filename)
	fs := mkfs(filename)
	mkdir(fs, "/A")
	mkdir(fs, "/b")
	mkfile(fs, "/testfile")
	mkfile(fs, "/b/sub")
}
func mkfs(name string) filesystem.FileSystem {
	size := int64(10 * 1024 * 1024)
	d, err := diskfs.Create(name, size, diskfs.Raw)
	if err != nil {
		fmt.Printf("Error creating disk: %v", err)
		os.Exit(1)
	}

	spec := disk.FilesystemSpec{
		Partition: 0,
		FSType:    filesystem.TypeFat32,
	}

	fs, err := d.CreateFilesystem(spec)
	if err != nil {
		panic(err)
	}
	return fs
}

func mkfile(fs filesystem.FileSystem, name string) {
	rw, err := fs.OpenFile(name, os.O_CREATE|os.O_RDWR)
	if err != nil {
		panic(err)
	}

	_, err = rw.Write([]byte("Hello World"))
	if err != nil {
		panic(err)
	}
}

func mkdir(fs filesystem.FileSystem, name string) {
	err := fs.Mkdir(name)
	if err != nil {
		panic(err)
	}
}

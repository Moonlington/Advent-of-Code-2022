package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

func newFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

func fileFromLine(line string) *File {
	sizeString, name, _ := strings.Cut(line, " ")
	size, err := strconv.Atoi(sizeString)
	if err != nil {
		panic(err)
	}
	return newFile(name, size)
}

type Directory struct {
	Name            string
	ParentDirectory *Directory
	Files           []*File
	Subdirectories  []*Directory
}

func newDirectory(name string) *Directory {
	return &Directory{
		Name:            name,
		ParentDirectory: nil,
		Files:           []*File{},
		Subdirectories:  []*Directory{},
	}
}

func (d *Directory) newSubDirectory(name string) *Directory {
	directory := &Directory{
		Name:            name,
		ParentDirectory: d,
		Files:           []*File{},
		Subdirectories:  []*Directory{},
	}
	d.Subdirectories = append(d.Subdirectories, directory)
	return directory
}

func (d *Directory) cd(directoryName string) *Directory {
	if directoryName == ".." {
		return d.ParentDirectory
	}
	for _, subdirectory := range d.Subdirectories {
		if subdirectory.Name == directoryName {
			return subdirectory
		}
	}
	return d.newSubDirectory(directoryName)
}

func (d *Directory) addFile(file *File) {
	d.Files = append(d.Files, file)
}

func (d *Directory) getSize() int {
	sumSize := 0
	for _, file := range d.Files {
		sumSize += file.Size
	}
	for _, subDirectory := range d.Subdirectories {
		sumSize += subDirectory.getSize()
	}
	return sumSize
}

func (d *Directory) forAllSubDirectories(function func(d *Directory)) {
	function(d)
	for _, subdirectory := range d.Subdirectories {
		subdirectory.forAllSubDirectories(function)
	}
}

var Filesystem *Directory = newDirectory("/")
var CurrentDirectory *Directory = Filesystem

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ ") {
			line = line[2:]
			command, argument, _ := strings.Cut(line, " ")
			switch command {
			case "cd":
				if argument == "/" {
					break
				}
				CurrentDirectory = CurrentDirectory.cd(argument)
			case "ls":
				break
			}
		} else {
			if strings.HasPrefix(line, "dir ") {
				CurrentDirectory.newSubDirectory(line[4:])
			} else {
				CurrentDirectory.addFile(fileFromLine(line))
			}
		}
	}

	part1Solution := 0
	summer := func(d *Directory) {
		size := d.getSize()
		if size <= 100000 {
			part1Solution += size
		}
	}
	Filesystem.forAllSubDirectories(summer)
	fmt.Println(part1Solution)

	totalSpace := 70000000
	neededSpace := 30000000
	unusedSpace := totalSpace - Filesystem.getSize()
	var smallestDirectoryToDelete *Directory
	finder := func(d *Directory) {
		size := d.getSize()
		if unusedSpace+size < neededSpace {
			return
		}
		if smallestDirectoryToDelete == nil || smallestDirectoryToDelete.getSize() > size {
			smallestDirectoryToDelete = d
		}
	}
	Filesystem.forAllSubDirectories(finder)
	part2Solution := smallestDirectoryToDelete.getSize()
	fmt.Println(part2Solution)
}

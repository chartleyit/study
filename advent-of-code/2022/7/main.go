package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Dir struct {
	name           string
	du             int
	parent         *Dir
	subdirectories []*Dir
	files          []*File
}

type File struct {
	name string
	size int
}

func readFile(file string) []string {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			fileLines = append(fileLines, line)
		}
	}

	readFile.Close()

	return fileLines
}

func isCommand(s string) bool {
	// Compile the regular expression
	// regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	regex := regexp.MustCompile(`^\$`)

	// if it is a command
	if regex.MatchString(s) {
		fmt.Printf("string matches: %s\n", s)
		cmdParts := strings.Split(s, " ")
		switch cmdParts[1] {
		case "cd":
			fmt.Println("cd")
		case "ls":
			fmt.Println("ls")
		default:
			fmt.Println("command unknown")
		}
		return true
	}
	return false
}

func CommandHandler(s string, fs *Dir) {
}

func treverseFs(d *Dir, ind string) (int, int) {
	var duAllSubdirs int
	var duTotal int
	var duMaxDir int
	var dk int

	maxDirSize := 100000
	duSubdir := make(map[int]int, len(d.subdirectories))

	if d.subdirectories == nil {
		return duAllSubdirs, duMaxDir
	}

	// this is the recursion point
	// a couple of conditions that can happen
	// - we have multiple directories whos sum exceeds 100000
	// - the sum of the current directory and all its subdirectories exceeds 100000

	for i, x := range d.subdirectories {
		duSubdir[i], dk = treverseFs(x, "|  "+ind)
		duMaxDir = duMaxDir + dk
		duAllSubdirs = duAllSubdirs + duSubdir[i]
	}
	if d.du < maxDirSize {
		duMaxDir = duMaxDir + d.du
	}
	// sum of pwd and subdirs
	duTotal = d.du + duAllSubdirs

	fmt.Printf("%v%+v %d %d\n", ind, d.name, d.du, duMaxDir)
	// for _, f := range d.files {
	// 	fmt.Printf("%v%+v\n", ind, f.name)
	// }

	return duTotal, duMaxDir
}

func main() {
	// input is a single line
	input := readFile("input.part1")
	// input := readFile("input.sample")

	// initialize file system
	fsRoot := &Dir{"/", 0, nil, []*Dir{}, []*File{}}
	fmt.Printf("Initialize fs: %+v", fsRoot)
	pwd := fsRoot

	for _, x := range input {
		fmt.Printf("current working directory is %+v\n", pwd)

		fmt.Println("input: ", x)

		// just do it all here
		cmdParts := strings.Split(x, " ")
		if cmdParts[1] == "cd" {
			// TODO things to handle that aren't handled here
			// * '-' previous directory
			// * 'a/b/c' multiple directories in one cd
			if cmdParts[2] == "/" {
				pwd = fsRoot
			} else if cmdParts[2] == ".." {
				// go up a level
				pwd = pwd.parent
			} else {
				// create new and set parent
				newDir := &Dir{cmdParts[2], 0, pwd, []*Dir{}, []*File{}}
				// set child
				pwd.subdirectories = append(pwd.subdirectories, newDir)
				fmt.Printf("Current dir: %+v\n", pwd)
				fmt.Printf("New dir: %+v\n", newDir)
				pwd = newDir
			}
		} else if cmdParts[1] == "ls" {

		} else if cmdParts[1] == "dir" {
			// newDir := &Dir{cmdParts[1], 0, pwd, []*Dir{}, []*File{}}
			// pwd.subdirectories = append(pwd.subdirectories, newDir)
		} else {
			fileSize, _ := strconv.Atoi(cmdParts[0])
			file := &File{cmdParts[1], fileSize}
			pwd.du = pwd.du + fileSize
			pwd.files = append(pwd.files, file)
			fmt.Printf("current dir: %v\nfiles: %+v\n", pwd.name, pwd.files)
		}

		// flow logic here; if it is a command we need to parse the next couple of lines
		// if command true we continue to parse until command true again
		// fmt.Printf("is %s a command? %+v\n", x, isCommand(x))
	}
	fmt.Printf("%+v\n", fsRoot)
	_, total := treverseFs(fsRoot, "")
	fmt.Printf("%v\n", total)
}

package utils

import (
	"fmt"
	"log"
	"os"
	"path"
)

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// WriteToFile creates a file and writes content to it
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}

func WriteToPath(filePath, fileName, content string) {
	// create file on path if exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0755)

		appPath, _ := os.Getwd()
		os.Mkdir(path.Join(appPath, filePath), 0755)

		fmt.Printf("\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, filePath, fileName), "\x1b[0m")

		WriteToFile(path.Join(appPath, filePath, fileName+".go"), content)
	} else {
		log.Println("File exist. " + fileName + " is not created")
	}
}

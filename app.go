package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func getFileStat(fileInfo fs.FileInfo) (string, uint) {
	modeString := fileInfo.Mode().String()
	fileSize := fileInfo.Size()
	return modeString, uint(fileSize)
}

func getDirectoryStat(dirInfo fs.FileInfo) string {
	modeString := dirInfo.Mode().String()
	return modeString
}

func writeStringToFile(f *os.File, data string) {
	_, err := f.WriteString(data)
	if err != nil {
		panic(err)
	}

}

func main() {
	var dirPath string
	var outName string

	if len(os.Args) != 3 {
		fmt.Println("Incorrect use of binary.")
		fmt.Println("The correct use is:")
		fmt.Println("\twalk-dir(.exe) [DIRECTORY_LOCATION] [OUTPUT_FILENAME]")
		return
	}

	dirPath = os.Args[1]
	outName = os.Args[2]

	var minimumFileSize uint = ^uint(0)
	var maximumFileSize uint = 0

	filePermsMap := make(map[string]int)
	dirPermsMap := make(map[string]int)

	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirPerm := getDirectoryStat(info)
			_, found := dirPermsMap[dirPerm]
			if !found {
				dirPermsMap[dirPerm] = 0
			}
		} else {
			filePerm, fileSize := getFileStat(info)
			_, found := filePermsMap[filePerm]
			if !found {
				filePermsMap[filePerm] = 0
			}
			if fileSize > maximumFileSize {
				maximumFileSize = fileSize
			}
			if fileSize < minimumFileSize {
				minimumFileSize = fileSize
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	f, err := os.Create(outName)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer f.Close()
	writeStringToFile(f, "Maximum Size of a File in Directory: "+fmt.Sprint(maximumFileSize/(1024*1024))+" MB\n")
	writeStringToFile(f, "Minimum Size of a File in Directory: "+fmt.Sprint(minimumFileSize/(1024*1024))+" MB\n\n")

	writeStringToFile(f, "All Directory Permissions encountered: \n")
	for perm := range dirPermsMap {
		writeStringToFile(f, perm+"\n")
	}
	writeStringToFile(f, "\n")
	writeStringToFile(f, "All Files Permissions encountered: \n")
	for perm := range filePermsMap {
		writeStringToFile(f, perm+"\n")
	}

	fmt.Println("Done!")
}

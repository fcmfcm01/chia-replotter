package utils

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE_LIST_FIILE = "file_list.txt"

func getFileList(dir string, prefix string, suffix string) ([]fs.FileInfo, int64) {
	dirInfo, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Expecting dir is not exists:" + dir)
			panic(err)
		}
	} else if !dirInfo.IsDir() {
		log.Fatalf("Input path:[%s] is not a directory.", dir)
		panic("Input path:[" + dir + "] is not a directory.")
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
		return nil, 0
	}
	var fileList []fs.FileInfo
	var size int64
	for _, file := range files {
		fileName := file.Name()
		if strings.HasPrefix(fileName, prefix) && strings.HasSuffix(fileName, suffix) {
			//log.Println(dir+"/"+fileName+" found")
			size += file.Size()
			fileList = append(fileList, file)
		}
	}
	return fileList, size
}

func RemoveOldFiles(oldDir string, newDir string, prefix string, suffix string) error {
	oldFiles, _ := getFileList(oldDir, prefix, suffix)

	newFiles, newFilesSize := getFileList(newDir, prefix, suffix)
	previousFileSize := getPreviousFileSize(newDir, newFilesSize)
	previousFiles := getPreviousFiles(newDir, newFiles)

	if newFilesSize > previousFileSize {
		log.Printf("new Files has been added:  \n")
		for i, newFile := range compareFiles(newFiles, previousFiles) {
			log.Printf("%d:\t%s \t%d", i, newFile.Name(), newFile.Size())
		}

		for _, oldFile := range oldFiles {
			if oldFile.Size() < (newFilesSize - previousFileSize) {
				log.Printf("removing old file:%s with size(%d) ...\n", oldFile.Name(), oldFile.Size())
				err := os.Remove(oldDir + "/" + oldFile.Name())
				if err != nil {
					log.Fatalln(oldDir+"/"+oldFile.Name()+" deletion failed: ", err)
				} else {
					log.Printf("removed old file:%s \n", oldFile.Name())
					writeFileList(newDir, newFiles)
					writeCurrentFileSize(newDir, newFilesSize)
				}
				break
			}
		}
	} else {
		log.Println("No changing...")
	}

	return nil
}

func compareFiles(newFiles []fs.FileInfo, oldFiles []fs.FileInfo) []fs.FileInfo {
	var diff []fs.FileInfo
	for _, srcFile := range newFiles {
		sameFileFind := false
		for _, destFile := range oldFiles {
			if srcFile.Name() == destFile.Name() {
				sameFileFind = true
				break
			}
		}
		if !sameFileFind {
			log.Println(srcFile.Name())
			diff = append(diff, srcFile)
		}
	}

	return diff
}

func getPreviousFiles(dirName string, files []fs.FileInfo) []fs.FileInfo {
	var oldFiles []fs.FileInfo
	filePath := dirName + "/file_list.txt"
	if fileExists(filePath) {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalln(err)
			return nil
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fileInfo, err := os.Stat(scanner.Text())
			if err != nil {
				log.Fatalln(err)
			} else {
				oldFiles = append(oldFiles, fileInfo)
			}
		}
	} else {
		log.Println("writing files to " + filePath)
		oldFiles = files
		writeFileList(dirName, oldFiles)
	}
	return oldFiles
}

func writeFileList(dir string, files []fs.FileInfo) {
	filePath := dir + "/" + FILE_LIST_FIILE
	os.Remove(filePath)
	file, _ := os.Create(filePath)
	bufferedWriter := bufio.NewWriter(file)

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, dir+"/"+file.Name())
	}
	log.Println("Files will be write to ["+filePath+"] :", strings.Join(fileNames, "\n"))
	_, _ = bufferedWriter.WriteString(strings.Join(fileNames, "\n"))
	bufferedWriter.Flush()
}

func getPreviousFileSize(dirName string, newFilesSize int64) int64 {
	filePath := dirName + "/file_size.txt"
	var previousFileSize int64
	if fileExists(filePath) {
		//log.Println(filePath + " exists")
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalln(err)
		} else {
			previousFileSize, err = strconv.ParseInt(string(content), 10, 64)
		}
	} else {
		log.Println(filePath + " not exists")
		previousFileSize = newFilesSize
		writeCurrentFileSize(dirName, previousFileSize)
	}
	return previousFileSize
}
func writeCurrentFileSize(dirName string, newFileSize int64) {
	filePath := dirName + "/file_size.txt"
	if fileExists(filePath) {
		os.Remove(filePath)
		os.Create(filePath)
	}
	ioutil.WriteFile(filePath, strconv.AppendInt([]byte(""), newFileSize, 10), 0644)

}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

package datareader

import (
	"io/ioutil"
)

func ReadDataFromFile(filePath string) (*GroupSchedule, error) {
	fileData, err := ioutil.ReadFile(filePath)
	// can file be opened?
	if err != nil {
		return nil, err
	}

	result, err := NewGroupSchedule(fileData)

	if err != nil {
		return nil, err
	}

	return result, err
}

func GetFileNameFromDir(dirPath string) ([]string, error) {
	var fileList []string
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fileList, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList, nil
}

package main

import "io/ioutil"

func main() {

	makeIndex(getFilesNames("files"))
}

func getFilesNames(path string) []string {
	files, _ := ioutil.ReadDir(path)

	stringFiles := []string{}

	for _, file := range files {
		stringFiles = append(stringFiles, path+"/"+file.Name())
	}
	return stringFiles
}

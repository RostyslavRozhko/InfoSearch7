package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ZoneInfo struct {
	documentFreq int
	title        bool
	body         bool
}

type TermInfo struct {
	freq        int
	postingList map[int]*ZoneInfo
}

func makeIndex(files []string) {
	dictionary := make(map[string]*TermInfo)

	lenFiles := len(files)
	for docID := 0; docID < lenFiles; docID++ {
		fd, err := os.Open(files[docID])
		check(err)
		scanner := bufio.NewScanner(fd)
		scanner.Split(ScanTerms)

		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			termInfo, ok := dictionary[token]
			if ok {
				termInfo.freq++
				zoneInfo, ok := termInfo.postingList[docID]
				if ok {
					zoneInfo.documentFreq++
				} else {
					termInfo.postingList[docID] = &ZoneInfo{1, false, true}
				}
			} else {
				termMap := make(map[int]*ZoneInfo)
				termMap[docID] = &ZoneInfo{1, false, true}
				dictionary[token] = &TermInfo{1, termMap}
			}
		}
	}
	print(dictionary)
}

func print(dictionary map[string]*TermInfo) {
	for key, value := range dictionary {
		fmt.Print(key, " ", value.freq, ": ")
		for docID, info := range value.postingList {
			fmt.Print(docID, " \"", info.documentFreq, info.body, info.title, "\" ")
		}
		fmt.Println("")
	}
}

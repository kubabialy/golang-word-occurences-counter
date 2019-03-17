// This app will parse srt and txt files
// count the word occurences and return an array of words with occurrences ordered by their occurences

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Provide absolute path for the file to analyse:")
	fpath, _ := reader.ReadString('\n')

	fpath = strings.TrimSuffix(fpath, "\n")

	ext := filepath.Ext(fpath)

	f, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
		return
	}

	switch ext {
	case ".txt":
		sortElements(parseFileContent(string(f)))
	case ".srt":
		sortElements(parseFileContent(string(f)))
	default:
		println("File type not supported. Provide .srt or .txt")
		return
	}
}

func parseFileContent(fileContents string) map[string]int {
	re := regexp.MustCompile(`(?m)[\pL_]`)
	wCol := strings.Split(string(fileContents), " ")
	wFound := make(map[string]int)
	for _, word := range wCol {
		if re.MatchString(word) {
			if val, found := wFound[word]; found {
				wFound[word] = val + 1
			} else {
				wFound[word] = 1
			}
		}
	}

	return wFound
}

// create map[int]string with key iterating from 0 to n and values being keys of wordsMap
// then bubble sort the wordsMap using previous map to find the values

func sortElements(wordsMap map[string]int) {
	type wo struct {
		Word       string
		Occurences int
	}
	var sortSlice []wo
	for k, v := range wordsMap {
		sortSlice = append(sortSlice, wo{k, v})
	}

	sort.Slice(sortSlice, func(i int, j int) bool {
		return sortSlice[i].Occurences > sortSlice[j].Occurences
	})

	for _, v := range sortSlice {
		fmt.Printf("%+v\n", v)
	}
}

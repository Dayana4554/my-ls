package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func R(path string) (map[string][]string, []string) {
	temporder := tempR(path)
	sortedMap := make(map[string][]string)
	var result []string
	var att []string

	for key := range temporder {
		att = append(att, key)
	}

	Sort(att)

	if lst {
		var last []string
		for key, value := range temporder {
			temp := reverseStringSlice(t(value, key))
			for _, r := range temp {
				lasttemp := getLastPathComponent(r)
				last = append(last, lasttemp)
			}
			sortedMap[key] = last
			last = nil
		}
		att = customSort(att)

	} else {
		for key, value := range temporder {
			sortedMap[key] = value
		}
	}

	if lsr {
		for key, values := range sortedMap {
			tempvalue := r(values)
			sortedMap[key] = tempvalue
		}
		result = reverseStringsWithSameSlashCount(att)
	} else {
		result = append(result, att...)
	}

	return sortedMap, result
}

func tempR(path string) map[string][]string {
	alldir := make(map[string][]string)

	var exploreDir func(string)
	exploreDir = func(dirPath string) {
		fileInfos, err := os.ReadDir(dirPath)
		if err != nil {
			fmt.Printf("Erreur de lecture de %s : %v\n", dirPath, err)
			return
		}

		contents := []string{}
		if lsa {
			f, _ := os.Open("..")
			defer f.Close()
			fi, _ := f.Stat()
			contents = append(contents, fi.Name())

			f2, _ := os.Open(".")
			defer f2.Close()
			fj, _ := f2.Stat()
			contents = append(contents, fj.Name())
			for _, fileInfo := range fileInfos {
				name := fileInfo.Name()
				contents = append(contents, name)

				if fileInfo.IsDir() {
					subdir := dirPath + "/" + name
					exploreDir(subdir)
				}
			}
		} else {
			for _, fileInfo := range fileInfos {
				if fileInfo.Name()[0] != '.' && !lsa {
					name := fileInfo.Name()
					contents = append(contents, name)

					if fileInfo.IsDir() {
						subdir := dirPath + "/" + name
						exploreDir(subdir)
					}
				}
			}
		}

		contents = Sort(contents)

		alldir[dirPath] = contents
	}

	exploreDir(path)

	return alldir
}

func reverseStringSlice(input []string) []string {
	length := len(input)
	reversed := make([]string, length)

	for i := 0; i < length; i++ {
		reversed[i] = input[length-i-1]
	}

	return reversed
}


func customSort(arr []string) []string {
	// Crée une carte pour regrouper les chaînes par nombre de "/".
	groups := make(map[int][]string)

	// Parcourez le tableau d'entrée et répartissez les chaînes en fonction du nombre de "/".
	for _, str := range arr {
		count := strings.Count(str, "/")
		groups[count] = append(groups[count], str)
	}

	// Trie chaque groupe de chaînes chronologiquement.
	for _, group := range groups {
		group, _ = sortByTimeRR(group)
	}
	result := CombineStringsFromMap(groups)

	return result
}

func CombineStringsFromMap(inputMap map[int][]string) []string {
	// Créez une tranche (slice) pour stocker toutes les chaînes.
	var combinedStrings []string

	// Créez une tranche pour stocker les clés (c'est-à-dire le nombre de "/") triées.
	var sortedKeys []int
	for key := range inputMap {
		sortedKeys = append(sortedKeys, key)
	}

	// Triez les clés en ordre croissant.
	bubbleSort(sortedKeys)

	// Parcourez les clés triées et ajoutez les chaînes correspondantes à la tranche combinée.
	for _, key := range sortedKeys {
		combinedStrings = append(combinedStrings, inputMap[key]...)
	}

	return combinedStrings
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Échangez les éléments si ils sont dans le mauvais ordre
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// Si aucun échange n'a eu lieu dans cette passe, le tableau est trié
		if !swapped {
			break
		}
	}
}

func sortByTimeRR(files []string) ([]string, error) {
	// Crée une carte pour stocker les fichiers et leurs horodatages
	fileTimes := make(map[string]time.Time)

	// Obtient les horodatages des fichiers
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return nil, err
		}
		fileTimes[file] = info.ModTime()
	}

	// Trie les fichiers en utilisant l'algorithme de tri par insertion
	for i := 1; i < len(files); i++ {
		j := i
		for j > 0 && fileTimes[files[j-1]].Before(fileTimes[files[j]]) {
			files[j-1], files[j] = files[j], files[j-1]
			j--
		}
	}

	return files, nil
}

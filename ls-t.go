package main

import (
	"os"
	"strings"
)

func t(infos []string, path string) []string {
	var array []string
	// pb avec le directorie "."
	if lsR {
		for _, r := range infos {
			array = append(array, path+"/"+r)
		}
		for i := 0; i < len(array)-1; i++ {
			for j := i + 1; j < len(array); j++ {
				fileInfoI, _ := os.Stat(array[i])
				fileInfoJ, _ := os.Stat(array[j])
				if fileInfoI.ModTime().Equal(fileInfoJ.ModTime()) {
					// Si les dates de modification sont égales, comparez les noms des fichiers
					if strings.ToLower(array[i]) > strings.ToLower(array[j]) {
						array[i], array[j] = array[j], array[i]
					}
				} else if fileInfoI.ModTime().Before(fileInfoJ.ModTime()) {
					array[i], array[j] = array[j], array[i]
				}
			}
		}
		var tab []string 
		if lsa && lsR && lsl && lsr && lst {
			for i := len(array)-1; i >= 0; i-- {
				tab = append(tab, array[i])
			}
		} else {
			for _, r := range array {
				tab = append(tab, r)
			}
		}
		return tab 
	} else {
		// fmt.Println(infos)
		array = append(array, infos...)

		for i := 0; i < len(array)-1; i++ {
			for j := i + 1; j < len(array); j++ {
				fileInfoI, _ := os.Stat(path + "/" + array[i])
				fileInfoJ, _ := os.Stat(path + "/" + array[j])
				 if fileInfoI.ModTime().Before(fileInfoJ.ModTime()) {
					array[i], array[j] = array[j], array[i]
				}
			}
		}

		// array, _ = sortByTime(array)
		return array
	}
}

func Sort(arr []string) []string {
	temp := false
	for _, r := range os.Args[1:] {
		if r[0] != '-' {
			if r == "/usr/bin" {
				temp = true
				break
			}
		}
	}
	if temp {
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				tempi := ""
				tempj := ""
				for _, r := range arr[i] {
					if (r >= 33 && r <= 47) || (r >= 58 && r <= 64) || (r >= 91 && r <= 96) || (r >= 123 && r <= 127) {
						// Ignore les caractères spéciaux
						continue
					}
					tempi += string(r)
				}
				for _, r := range arr[j] {
					if (r >= 33 && r <= 45) || (r == 47) || (r >= 58 && r <= 64) || (r >= 91 && r <= 96) || (r >= 123 && r <= 127) {
						// Ignore les caractères spéciaux
						continue
					}
					tempj += string(r)
				}
				if strings.ToLower(tempi) > strings.ToLower(tempj) {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	} else {
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				if strings.ToLower(arr[i]) > strings.ToLower(arr[j]) {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}
	return arr
}

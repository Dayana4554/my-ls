package main

import (
	"fmt"
	"os"
	"strings"
)

func a(existingPaths []string) [][]string {
	var result [][]string

	if len(existingPaths) != 0 {
		for _, r := range existingPaths {
			var temp []string
			path, _ := getAbsolutePath(r)
			/*if path[len(path)-1] == '/' {
				path = path[:len(path)-1]
			}*/
			temp = append(temp, ".")
			// uppath := getLastDirectory(path)
			temp = append(temp, "..")

			files := listContents(path)
			for _, r := range files {
				temp = append(temp, getFileNameFromPath(r))
			}

			temp = Sort(temp)

			result = append(result, temp)
			temp = nil

		}
	}

	return result
}

func getAbsolutePath(directoryPath string) (string, error) {
	if strings.HasPrefix(directoryPath, "/") {
		// Le chemin est déjà absolu, le renvoyer tel quel
		return directoryPath, nil
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	absolutePath := currentDir + "/" + directoryPath
	return absolutePath, nil
}

func listContents(path string) []string {
	var contents []string

	fileInfos, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Erreur de lecture de %s : %v\n", path, err)
		return contents
	}

	for _, fileInfo := range fileInfos {
		if path == "." {
			if !lsa {
				if fileInfo.Name()[0] != '.' {
					contents = append(contents, fileInfo.Name())
				}
			} else {
				contents = append(contents, fileInfo.Name())
			}
		} else {
			if !lsa {
				if fileInfo.Name()[0] != '.' {
					contents = append(contents, "./"+path+"/"+fileInfo.Name())
				}
			} else {
				contents = append(contents, "./"+path+"/"+fileInfo.Name())
			}
		}
	}

	return contents
}

func getLastDirectory(path string) string {
	segments := strings.Split(path, "/")
	if len(segments) <= 1 {
		return ""
	}

	lastSegment := segments[len(segments)-1]
	return path[:strings.LastIndex(path, lastSegment)-1]
}

func getFileNameFromPath(path string) string {
	segments := strings.Split(path, "/")
	if len(segments) > 0 {
		return segments[len(segments)-1]
	}
	return ""
}

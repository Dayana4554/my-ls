package main

import (
	"strings"
)

func r(array []string) []string {
	count := 0
	for _, r := range array {
		if !strings.Contains(r, ".") || r == "." || r == ".." {
			count++
		}
	}

	if count == len(array) {
		return reverseStringsWithSameSlashCount(array)
	} else {
		reverseArray(array)
		return array
	}
}

func reverseArray(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func countSlash(s string) int {
	count := 0
	for _, char := range s {
		if char == '/' {
			count++
		}
	}
	return count
}

func reverseStringsWithSameSlashCount(arr []string) []string {
	counts := make(map[int][]string)
	var zeroSlashCount []string

	for _, s := range arr {
		slashCount := countSlash(s)
		if slashCount == 0 {
			zeroSlashCount = append(zeroSlashCount, s)
		} else {
			counts[slashCount] = append(counts[slashCount], s)
		}
	}

	result := make([]string, 0)

	for _, s := range zeroSlashCount {
		result = append(result, s)
	}

	for count := 1; count <= len(arr); count++ {
		stringsWithSameCount, ok := counts[count]
		if !ok {
			continue
		}
		for i := len(stringsWithSameCount) - 1; i >= 0; i-- {
			result = append(result, stringsWithSameCount[i])
		}
	}

	return result
}

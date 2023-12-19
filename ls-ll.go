package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

//------------- command ls -l -------------------------------------


func l(path string, files []string) ([][]string, int64) {
	var total int64
	var result [][]string
	verif := false

	if path == "/dev" {
		temp, temptotal := lfile(path+"/"+files[0], verif)
		result = append(result, temp)
		total += temptotal
		temp, temptotal = lfile(path+"/"+files[1], verif)
		result = append(result, temp)
		total += temptotal
		dir, _ := os.ReadDir(path)
		for _, entry := range dir {
			array, _ := lfile2(entry, path)
			result = append(result, array)
		}
	} else {
		if path != "" {
			verif = true
			for _, file := range files {
				if file == "'['" {
					temp, totemp := lfile(path+"/[", verif)
					result = append(result, temp)
					total += totemp
				} else {
					temp, totemp := lfile(path+"/"+file, verif)
					result = append(result, temp)
					total += totemp
				}
			}
		} else {
			for _, file := range files {
				temp, _ := lfile(file, verif)
				result = append(result, temp)
			}
			total = 0
		}
	}

	total = total / 2

	return result, total
}

func lfile2(entry fs.DirEntry, path string) ([]string, int64) {
	// Parcourir les informations sur les fichiers
	var result []string
	var total int64

	f1, err := entry.Info()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	sys := f1.Sys().(*syscall.Stat_t)
	user, _ := user.LookupId(strconv.Itoa(int(sys.Uid)))
	username := user.Username

	nlink := uint64(0)
	if sys := f1.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			nlink = uint64(stat.Nlink)
		}
	}

	modTime := f1.ModTime()
	formattedDate := modTime.Format("2 Jan 06 15:04 MST")

	newdate := strings.Split(formattedDate, " ")

	var time string

	if strings.Contains(newdate[1], "Jan") {
		time += "janv."
	} else if strings.Contains(newdate[1], "Feb") {
		time += "févr."
	} else if strings.Contains(newdate[1], "Mar") {
		time += "mars"
	} else if strings.Contains(newdate[1], "Apr") {
		time += "avri."
	} else if strings.Contains(newdate[1], "May") {
		time += "mai"
	} else if strings.Contains(newdate[1], "Jun") {
		time += "juin"
	} else if strings.Contains(newdate[1], "Jul") {
		time += "juil."
	} else if strings.Contains(newdate[1], "Aug") {
		time += "août"
	} else if strings.Contains(newdate[1], "Sep") {
		time += "sept."
	} else if strings.Contains(newdate[1], "Oct") {
		time += "octo."
	} else if strings.Contains(newdate[1], "Nov") {
		time += "nove."
	} else if strings.Contains(newdate[1], "Dec") {
		time += "déce."
	}

	time += " " + newdate[0] + " " + newdate[3]

	var info fs.FileInfo
	if path != "" {
		info, _ = os.Lstat(path + "/" + entry.Name())
	} else {
		info, _ = os.Lstat(entry.Name())
	}

	sys = info.Sys().(*syscall.Stat_t)

	modeString := f1.Mode().String()
	sizeString := strconv.FormatInt(int64(f1.Size()), 10)

	myString := strconv.FormatUint(nlink, 10)

	total += int64(sys.Blocks)

	if sizeString == "0" {
		sizeString = strconv.FormatInt(int64(sys.Rdev/256), 10) + ", " + strconv.FormatInt(int64(sys.Rdev%256), 10)
	}

	if modeString[0] == 'D' {
		modeString = modeString[1:]
	} else if modeString[0] == 'L' {
		modeString = "l" + modeString[1:]
	}
	result = append(result, modeString)
	result = append(result, myString)
	result = append(result, username)
	result = append(result, username)
	result = append(result, sizeString)
	result = append(result, time)

	if info.Mode()&os.ModeSymlink != 0 {
		realPath, err := os.Readlink(path + "/" + entry.Name())
		if err != nil {
			fmt.Println(err)
		}

		result = append(result, entry.Name()+" -> "+realPath)
	} else {
		result = append(result, entry.Name())
	}
	return result, total
}

func lfile(paths string, pathverif bool) ([]string, int64) {
	var result []string
	var total int64
	f, _ := os.Open(paths)
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	temp := f.Name()

	f1, err := os.Stat(temp)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	nlink := uint64(0)
	if sys := f1.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			nlink = uint64(stat.Nlink)
		}
	}

	sys := fi.Sys().(*syscall.Stat_t)
	user, _ := user.LookupId(strconv.Itoa(int(sys.Uid)))

	info, err := os.Lstat(paths)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var sizeString string
	var time string
	var modeString string
	if strings.Contains(paths, "/usr/bin") {
		if info.Mode()&os.ModeSymlink != 0 {
			sizeString = strconv.FormatInt(info.Size(), 10)
			modTime := info.ModTime()
			formattedDate := modTime.Format("2 Jan 06 15:04 MST")

			newdate := strings.Split(formattedDate, " ")

			if strings.Contains(newdate[1], "Jan") {
				time += "janv."
			} else if strings.Contains(newdate[1], "Feb") {
				time += "févr."
			} else if strings.Contains(newdate[1], "Mar") {
				time += "mars"
			} else if strings.Contains(newdate[1], "Apr") {
				time += "avri."
			} else if strings.Contains(newdate[1], "May") {
				time += "mai"
			} else if strings.Contains(newdate[1], "Jun") {
				time += "juin"
			} else if strings.Contains(newdate[1], "Jul") {
				time += "juil."
			} else if strings.Contains(newdate[1], "Aug") {
				time += "août"
			} else if strings.Contains(newdate[1], "Sep") {
				time += "sept."
			} else if strings.Contains(newdate[1], "Oct") {
				time += "octo."
			} else if strings.Contains(newdate[1], "Nov") {
				time += "nove."
			} else if strings.Contains(newdate[1], "Dec") {
				time += "déce."
			}

			time += " " + newdate[0] + " " + newdate[3]
			modeString = info.Mode().String()
		} else {
			sizeString = strconv.FormatInt(f1.Size(), 10)

			modTime := fi.ModTime()
			formattedDate := modTime.Format("Jan _2 2006")

			time = formattedDate

			modeString = info.Mode().String()
		}
	} else {
		sizeString = strconv.FormatInt(info.Size(), 10)
		modTime := info.ModTime()
		formattedDate := modTime.Format("2 Jan 06 15:04 MST")

		newdate := strings.Split(formattedDate, " ")

		if strings.Contains(newdate[1], "Jan") {
			time += "janv."
		} else if strings.Contains(newdate[1], "Feb") {
			time += "févr."
		} else if strings.Contains(newdate[1], "Mar") {
			time += "mars"
		} else if strings.Contains(newdate[1], "Apr") {
			time += "avri."
		} else if strings.Contains(newdate[1], "May") {
			time += "mai"
		} else if strings.Contains(newdate[1], "Jun") {
			time += "juin"
		} else if strings.Contains(newdate[1], "Jul") {
			time += "juil."
		} else if strings.Contains(newdate[1], "Aug") {
			time += "août"
		} else if strings.Contains(newdate[1], "Sep") {
			time += "sept."
		} else if strings.Contains(newdate[1], "Oct") {
			time += "octo."
		} else if strings.Contains(newdate[1], "Nov") {
			time += "nove."
		} else if strings.Contains(newdate[1], "Dec") {
			time += "déce."
		}

		time += " " + newdate[0] + " " + newdate[3]
		modeString = info.Mode().String()
	}

	if modeString[0] == 'L' {
		modeString = "l" + modeString[1:]
	}

	var myString string
	rep, _ := isSymbolicLinkToDir(paths)
	if rep && paths[len(paths)-1] != '/' {
		nlink := uint64(0)
		if sys := info.Sys(); sys != nil {
			if stat, ok := sys.(*syscall.Stat_t); ok {
				nlink = uint64(stat.Nlink)
			}
		}
		myString = strconv.FormatUint(nlink, 10)
	} else {
		myString = strconv.FormatUint(nlink, 10)
	}

	result = append(result, modeString)
	result = append(result, myString)
	result = append(result, user.Username)
	result = append(result, user.Username)
	result = append(result, sizeString)
	result = append(result, time)

	if info.Mode()&os.ModeSymlink != 0 {
		realPath, err := os.Readlink(paths)
		if err != nil {
			fmt.Println(err)
		}

		if pathverif {
			result = append(result, getLastPathComponent(f.Name())+" -> "+realPath)
		} else {
			result = append(result, paths+" -> "+realPath)
		}
	} else {
		sys := fi.Sys().(*syscall.Stat_t)
		total += sys.Blocks
		if pathverif {
			if f.Name() == "[" {
				result = append(result, "'['")
			} else {
				result = append(result, getLastPathComponent(f.Name()))
			}
		} else {
			if f.Name() == "[" {
				result = append(result, "'['")
			} else {
				if f.Name() == "/dev/." || f.Name() == "/dev/.." {
					result = append(result, getLastPathComponent(f.Name()))
				} else {
					result = append(result, paths)
				}
			}
		}
	}

	return result, total
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	lsl bool
	lsR bool
	lsa bool
	lsr bool
	lst bool
)

func main() {
	//-------------------------------verifie les flags-------------------------------------------------------------------------------------
	lsl = false
	lsR = false
	lsa = false
	lsr = false
	lst = false

	args := os.Args[1:]
	//----------------------------------gere les erreurs-----------------------------------------------------------------------------------
	if len(args) == 0 {
		fmt.Println("Aucune commande")
		os.Exit(1)
	}
	if args[0] != "ls" {
		fmt.Println("Commande invalide")
		os.Exit(1)
	}

	//------------------------------- tableau des fichiers/repertoires---------------------------------------------------------------------
	existingPaths := []string{}

	for _, r := range args[1:] {
		if r[0] == '-' {
			if len(r) != 1 {
				// Gerer les flags
				if strings.Contains(r, "a") {
					lsa = true
				}
				if strings.Contains(r, "l") {
					lsl = true
				}
				if strings.Contains(r, "R") {
					lsR = true
				}
				if strings.Contains(r, "r") {
					lsr = true
				}
				if strings.Contains(r, "t") {
					lst = true
				}
			} else {
				existingPaths = append(existingPaths, r)
			}
		} else {
			// Vérifier si le répertoire ou le fichier existe
			if _, err := os.Stat(r); err == nil {
				existingPaths = append(existingPaths, r)
			} else if os.IsNotExist(err) {
				fmt.Printf("%s n'existe pas\n", r)
				os.Exit(0)
			} else {
				fmt.Printf("Erreur lors de la vérification de %s : %v\n", r, err)
				os.Exit(0)
			}
		}
	}
	if len(existingPaths) == 0 {
		existingPaths = append(existingPaths, ".")
	}

	//------------------------------------------------je m occupe des flags et j applique----------------------------------------------------
	var files []string
	var directories []string
	for _, r := range existingPaths {
		test, _ := os.Stat(r)
		if test.IsDir() {
			directories = append(directories, r)
		} else {
			files = append(files, r)
		}
	}
	if len(files) > 1 {
		files = Sort(files)
	}

	if len(files) > 1 {
		directories = Sort(directories)
	}

	//---------------------------pas de flag----------------------------------------------

	if !lsa && !lsR && !lst && !lsr && !lsl {
		final := ""
		var array []string
		for _, r := range files {
			final += r
			final += "  "
		}

		if len(files) != 0 && len(directories) != 0 {
			final += "\n\n"
		}

		for _, r := range directories {
			if len(files) != 0 && len(directories) == 1 || len(directories) != 1 && len(files) == 0 || len(directories) != 1 && len(files) != 0 {
				final += r + ":\n"
			}
			temp, _ := os.Open(r)
			defer temp.Close()
			f, _ := temp.Readdir(-1)
			for _, k := range f {
				nom := k.Name()
				if nom[0] != '.' {
					array = append(array, k.Name())
				}
			}
			array = Sort(array)

			for _, l := range array {
				final += l
				final += "  "
			}
			if len(directories) != 1 && len(files) != 0 {
				if r != directories[len(directories)-1] {
					final += "\n\n"
				}
			}

			array = nil

		}
		fmt.Println(final)
		os.Exit(0)
	}

	if lsl {
		if lsR {
			final := ""

			if len(files) != 0 {
				if lst {
					files = t(files, "")
				}

				if lsr {
					files = r(files)
				}

				a, _ := l("", files)
				for i, k := range a {
					for _, r := range k {
						final += r + " "
					}
					if i != len(a)-1 {
						final += "\n"
					}
				}
			}

			if len(files) != 0 && len(directories) != 0 {
				final += "\n\n"
			}

			for _, r := range directories {
				dirPath := r
				a, result := R(dirPath)
				for i, r := range result {
					if len(files) != 0 && len(directories) != 1 || len(files) != 0 && len(directories) == 1 || len(files) == 0 && len(directories) != 1 || lsR {
						if strings.Contains(r, "////") {
							temp := strings.Split(r, "////")
							newstr := ""
							for _, r := range temp {
								newstr += r
								if r != temp[len(temp)-1] {
									newstr += "/"
								}
							}
							final += newstr + ":\n"
						} else {
							final += r + ":\n"
						}
					}
					affiche, total := l(r, a[r])
					str := strconv.FormatInt(total, 10)
					if total != 0 {
						if strings.Contains(r, "///") {
							final += "total " + str + "\n"
						} else {
							final += "total: " + str + "\n"
						}
					}
					for j, r := range affiche {
						for _, k := range r {
							final += k + " "
						}
						if i != len(result)-1 {
							final += "\n"
						} else {
							if j != len(affiche)-1 {
								final += "\n"
							}
						}
					}
					if i != len(directories)-1 {
						final += "\n"
					}
				}
			}
			fmt.Println(final)
			os.Exit(0)
		} else {
			final := ""

			if len(files) != 0 {
				if lst {
					files = t(files, "")
				}

				if lsr {
					files = r(files)
				}
				a, _ := l("", files)
				for i, k := range a {
					for _, r := range k {
						final += r + " "
					}
					if i != len(a)-1 {
						final += "\n"
					}
				}
			}

			if len(files) != 0 && len(directories) != 0 {
				final += "\n\n"
			}

			directories = Sort(directories)

			if lst {
				directories = t(directories, "")
			}

			if lsr {
				reverseArray(directories)
			}
			var contents []string
			for i, r := range directories {
				rep, _ := isSymbolicLinkToDir(r)
				if rep && r[len(r)-1] != '/' {
					temp := getSymlinkInfo(r)
					for _, k := range temp {
						final += k + " "
					}
				} else {
					if len(files) != 0 && len(directories) != 1 || len(files) != 0 && len(directories) == 1 || len(files) == 0 && len(directories) != 1 || lsR {
						final += r + ":\n"
					}
					fileInfos, _ := os.ReadDir(r)
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
						}
					} else {
						for _, fileInfo := range fileInfos {
							if fileInfo.Name()[0] != '.' && !lsa {
								name := fileInfo.Name()
								if strings.Contains(name, "[") {
									name = "'" + name + "'"
								}
								contents = append(contents, name)
							}
						}
					}
					contents = Sort(contents)

					if lst {
						contents = t(contents, r)
					}

					if lsr {
						reverseArray(contents)
					}

					affiche, total := l(r, contents)
					contents = nil

					str := strconv.FormatInt(total, 10)
					if len(files) != 0 && len(directories) != 1 || len(files) != 0 && len(directories) == 1 || len(files) == 0 && len(directories) != 1 || lsR {
						final += "total: " + str + "\n"
					} else {
						final += "total " + str + "\n"
					}
					for j, r := range affiche {
						for _, k := range r {
							final += k + " "
						}
						if i != len(directories)-1 {
							final += "\n"
						} else {
							if j != len(affiche)-1 {
								final += "\n"
							}
						}
					}

					if i != len(directories)-1 {
						final += "\n"
					}
				}
			}

			fmt.Println(final)
			os.Exit(0)
		}
	}

	//---------------------------flag -R-----------------------------------------------------------------------------------------------------

	if lsR {
		var result string
		if len(files) != 0 {
			var final []string
			var temp []string
			if lst {
				essais := t(files, "")
				temp = essais
			} else {
				files = Sort(files)
				temp = files
			}

			if lsr {
				final = r(temp)
				temp = final
			}
			for _, r := range temp {
				result += r
				result += "  "
			}
		}
		// var lastresult string
		if len(directories) != 0 && len(files) != 0 {
			result += "\n"
			result += "\n"
		}

		var resultat []string
		var tempRR map[string][]string
		// j en suis la
		for _, r := range directories {
			// var tempsort []string
			tempRR, resultat = R(r)
			for _, r := range resultat {
				for key, value := range tempRR {
					if r == key {
						result += string(key) + ":\n"
						for _, r := range value {
							result += r + "  "
						}

						if key != resultat[len(resultat)-1] {
							result += "\n"
							result += "\n"
						}
					}
				}
			}
		}
		fmt.Println(result)
		os.Exit(0)
	}

	//-----------------------------------------------flag -a-------------------------------------------------------------------------------

	if lsa {
		//--------------------------------------------------------------------------------
		if lsr {
			result := ""
			if len(files) != 0 {
				files = Sort(files)

				for _, r := range files {
					result += r
					result += "  "
				}
				if len(directories) != 0 {
					result += "\n"
					result += "\n"
				}
			}

			if len(directories) != 0 {
				directories = Sort(directories)

				temp := a(directories)

				if len(directories) > 1 && len(files) != 0 || len(directories) > 1 && len(files) == 0 {
					for i, r := range temp {
						result += directories[i] + ":"
						result += "\n"
						for j := len(r) - 1; j > 1; j-- {
							result += getFileNameFromPath(r[j])
							result += "  "
						}
						result += "..  .  "

						if i != len(temp)-1 {
							result += "\n"
							result += "\n"
						}
					}
				} else if len(directories) == 0 {
					for j := len(temp[0]) - 1; j > 1; j-- {
						result += getFileNameFromPath(temp[0][j])
						result += "  "
					}
					result += "..  .  "
				}
			}
			fmt.Println(result)
			os.Exit(0)

			//-----------------------------------------------------------------------------------------

			if lst {
				//------------------------------------------------------------------------------------------
			} else {
			}

			//-----------------------------------------------------------------------------------
		} else {
			result := ""

			if len(files) != 0 {
				files = Sort(files)

				for _, r := range files {
					result += r
					result += "  "
				}

				if len(directories) != 0 {
					fmt.Println()
				}
			}

			if len(directories) != 0 {
				if len(files) != 0 {
					fmt.Println()
					fmt.Println()
				}
			}

			directories = Sort(directories)

			temp := a(directories)

			if len(directories) > 1 && len(files) != 0 || len(directories) > 1 && len(files) == 0 {
				for i, r := range temp {
					result += directories[i] + ":"
					result += "\n"
					// result += ".  ..  "
					for j := 0; j < len(r); j++ {
						result += r[j]
						result += "  "
					}

					if i != len(temp)-1 {
						result += "\n"
						result += "\n"
					}
				}
			} else {
				// result += ".  ..  "
				for j := 0; j < len(temp[0]); j++ {
					result += temp[0][j]
					result += "  "
				}
			}
			fmt.Println(result)
		}
		os.Exit(0)
	}

	//----------------------------------------------flag -r--------------------------------------------------------------------------------
	if lsr { // fini
		if lst {
			var final string
			if len(files) != 0 {
				tempresult := t(files, "")
				result := r(tempresult)

				for _, r := range result {
					final += r
					final += "  "
				}

				if len(directories) != 0 {
					final += "\n\n"
				}
			}

			if len(directories) != 0 {
				if len(directories) == 1 && len(files) == 0 {
					temp := listContents(directories[0])
					arrayt := t(temp, directories[0])
					result := r(arrayt)

					for _, r := range result {
						final += r
						final += "  "
					}
				} else {
					for _, k := range directories {
						final += k + ":\n"
						temp := listContents(k)
						arrayt := t(temp, k)
						result := r(arrayt)
						for _, r := range result {
							if k == directories[len(directories)-1] && r == result[len(result)-1] {
								final += r
							} else {
								final += r + "  "
							}
						}
					}
				}
			}
			os.Exit(0)
		} else {
			final := ""
			if len(files) != 0 {
				result := r(files)

				for _, r := range result {
					final += r
					final += "  "
				}

				if len(directories) != 0 {
					final += "\n\n"
				}
			}

			if len(directories) != 0 {
				if len(directories) == 1 && len(files) == 0 {
					temp := listContents(directories[0])
					temp = Sort(temp)
					result := r(temp)

					for _, r := range result {
						final += r
						final += "  "
					}
				} else {
					for _, k := range directories {
						final += k + ":\n"
						temp := listContents(k)
						result := r(temp)
						for _, r := range result {
							if k == directories[len(directories)-1] && r == result[len(result)-1] {
								final += r
							} else {
								final += r + "  "
							}
						}
					}
				}
			}
			fmt.Println(final)
			os.Exit(0)
		}
	}

	//---------------------------------------------flag -t--------------------------------------------------------------------------------

	if lst { // fini
		final := ""

		if len(files) != 0 {
			result := t(files, "")

			for _, r := range result {
				final += r
				final += "  "
			}

			if len(directories) != 0 {
				final += "\n\n"
			}
		}

		if len(directories) != 0 {
			if len(directories) == 1 && len(files) == 0 {
				temp := listContents(directories[0])
				result := t(temp, directories[0])

				for _, r := range result {
					if strings.Contains(r, "/") {
						final += getLastPathComponent(r) + "  "
					} else {
						final += r + "  "
					}
				}
			} else {
				for _, k := range directories {
					final += k + ":\n"
					temp := listContents(k)
					result := t(temp, k)
					for _, r := range result {
						if k == directories[len(directories)-1] && r == result[len(result)-1] {
							if strings.Contains(r, "/") {
								final += getLastPathComponent(r)
							} else {
								final += r
							}
						} else {
							if strings.Contains(r, "/") {
								final += getLastPathComponent(r) + "  "
							} else {
								final += r + "  "
							}
						}
					}
				}
			}
		}

		fmt.Println(final)
		os.Exit(0)
	}
}

//-------------------------------------------fonctions que j appelles dans mon code-----------------------------------

func getSymlinkInfo(path string) []string {
	result, _ := lfile(path, false)

	return result
}

func getLastPathComponent(path string) string {
	components := strings.Split(path, "/")
	if len(components) > 0 {
		return components[len(components)-1]
	}
	return ""
}

func reverseAlphabeticalSort(names []string) {
	for i := 0; i < len(names)/2; i++ {
		j := len(names) - 1 - i
		names[i], names[j] = names[j], names[i]
	}
}

func isSymbolicLinkToDir(path string) (bool, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	// Vérifiez si le fichier est un lien symbolique
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		// Obtenez les informations sur le lien symbolique
		linkInfo, err := os.Stat(path)
		if err != nil {
			return false, err
		}

		// Vérifiez si le lien symbolique pointe vers un dossier
		if linkInfo.IsDir() {
			return true, nil
		}
	}

	return false, nil
}

func GetUpperPath(path string) string {
	pathInRune := []rune(path)
	var dashCounter int = 0

	for k := 0; k < len(pathInRune); k++ {
		if pathInRune[k] == '/' {
			dashCounter++
		}
	}

	if dashCounter == 1 {
		return "/"
	}

	for k := 0; k < len(path); k++ {
		if path[len(path)-1:] == "/" {
			path = path[:len(path)-1]
			break
		} else {
			path = path[:len(path)-1]
		}
	}

	return path
}

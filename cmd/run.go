package cmd

import (
	"fmt"
	"log"
	"os"
	"sortex/sortex"

	"github.com/spf13/cobra"
)

var currentDir bool
var determinedPath string
var internalDirs bool
var resultPath string
var moveWithoutEx bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&currentDir, "current", "c", false, "sort files in current dir")
	rootCmd.PersistentFlags().StringVarP(&determinedPath, "determined", "d", "", "select determined dir to sort")
	rootCmd.PersistentFlags().BoolVarP(&internalDirs, "internal-dirs", "i", false, "sort files in all internal dirs")
	rootCmd.PersistentFlags().StringVarP(&resultPath, "save", "s", "", "select the path where files will be sorted")
	rootCmd.PersistentFlags().BoolVarP(&moveWithoutEx, "move-no-data", "m", true, "should files without exif data be transfered?")
}

func sort(cmd *cobra.Command, args []string) {
	var path string
	var fileTypes []string = []string{"all"}

	if currentDir && determinedPath == "" {
		fmt.Println("[OK] sortex -c")
		p, err := os.Getwd()
		if err != nil {
			log.Fatal("Error while getting current dirrectory. Use -d")
		}
		path = p
	} else if !currentDir && determinedPath != "" {
		fmt.Println("[OK] sortex -d \"" + determinedPath + "\"")
		path = determinedPath
	} else if currentDir && determinedPath != "" {
		fmt.Println("[CANNOT USE] sortex -c -d \"" + determinedPath + "\"")
	} else {
		log.Fatal("[CANNOT USE] sortex -h --help")
	}

	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var backupData []sortex.BackupJSON
	var movePath = path

	for _, value := range dir {
		fullPath := fmt.Sprintf("%v/%v", path, value.Name())
		fmt.Println("full path", fullPath)

		if sortex.CheckTypes(fullPath, fileTypes) || fileTypes[0] == "all" {
			fmt.Println("--------------------------------")
			if internalDirs && sortex.IsItDir(fullPath) && fileTypes[0] == "all" {
				fmt.Println("path", fullPath, "is skipped because its directory ")
				continue
			}

			data, err := sortex.GetData(fullPath)
			if err != nil {
				log.Println(err)
				if !moveWithoutEx {
					continue
				}
			}
			fmt.Println("exif data:", data)

			if resultPath != "" {
				movePath = resultPath
			}
			newDir, err := sortex.CreateFolder(&data, movePath)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("newDir:", newDir)

			err = sortex.MoveToDir(fullPath, newDir+"/"+value.Name())
			if err != nil {
				log.Println(err)
			}

			backupData = append(backupData, sortex.BackupJSON{
				OldPath: fullPath,
				NewPath: newDir + "/" + value.Name(),
			})
		} else {
			fmt.Println("path", fullPath, "is skipped because its type is not in the list")
		}
		sortex.CreateBackup(path, backupData)
	}
}

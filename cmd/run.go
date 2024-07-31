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
var restoreBackup string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&currentDir, "current", "c", false, "(1)sort files in current dir")
	rootCmd.PersistentFlags().StringVarP(&determinedPath, "determined", "d", "", "(2)sort files in specific dir")
	//rootCmd.PersistentFlags().BoolVarP(&internalDirs, "internal-dirs", "i", false, "(1)(2)sort files in all internal dirs")
	rootCmd.PersistentFlags().StringVarP(&resultPath, "save", "s", "", "(1)(2)select the path where the files will be sorted")
	rootCmd.PersistentFlags().BoolVarP(&moveWithoutEx, "move-no-data", "m", true, "(1)(2)should files without data be transferred? (true by default)")
	rootCmd.PersistentFlags().StringVarP(&restoreBackup, "restore-backup", "b", "", "(0)undo changes if something went wrong")
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
	} else if restoreBackup != "" {
		err := sortex.RevertChanges(restoreBackup)
		if err != nil {
			log.Fatal("[ERROR] -", err)
		}
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
	}
	sortex.CreateBackup(movePath, backupData)
}

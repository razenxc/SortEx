package main

import (
	"SortFilesWithEXIFdata/sortfiles"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Path to dir with files to sort
	input, err := readString("Enter full path to dir. with files: ")
	if err != nil {
		log.Fatal(err)
	}
	path := input

	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// File types to sort
	input, err = readString("Enter files types to sort separated by spaces \".jpeg .png .jpg\" or \"all\": ")
	if err != nil {
		log.Fatal(err)
	}
	fileTypes := strings.Split(input, " ")

	// Ignore directories or not
	var ignoreDirs bool
	if fileTypes[0] == "all" {

		input, err = readString("Do you want ignore directories? [y/n]: ")
		if err != nil {
			log.Fatal(err)
		}

		switch input {
		case "y":
			ignoreDirs = true
		case "n":
			ignoreDirs = false
		default:
			fmt.Println("Invalid input. 'y' or 'n'.")
		}
	}

	for _, value := range dir {
		fullPath := fmt.Sprintf("%v/%v", path, value.Name())
		fmt.Println("full path", fullPath)

		if sortfiles.CheckTypes(fullPath, fileTypes) || fileTypes[0] == "all" {
			fmt.Println("--------------------------------")
			if ignoreDirs && sortfiles.IsItDir(fullPath) && fileTypes[0] == "all" {
				fmt.Println("path", fullPath, "is skipped because its directory ")
				continue
			}

			data, err := sortfiles.GetData(fullPath)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("exif data:", data)

			newDir, err := sortfiles.CreateFolder(&data, path)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("newDir:", newDir)

			err = sortfiles.MoveToFolder(fullPath, newDir+"/"+value.Name())
			if err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println("path", fullPath, "is skipped because its type is not in the list")
		}
	}
}

func readString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

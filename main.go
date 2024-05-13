package main

import (
	"EXIFphotoSorter/sortfiles"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path, err := readString("Enter full path to dir. with files: ")
	if err != nil {
		log.Fatal(err)
	}

	fileTypesStr, err := readString("Enter files types to sort separated by spaces \".jpeg .png .jpg\" or \"all\": ")
	if err != nil {
		log.Fatal(err)
	}
	fileTypes := strings.Split(fileTypesStr, " ")

	dir, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range dir {
		fullPath := fmt.Sprintf("%v/%v", path, value.Name())
		fmt.Println("full path", fullPath)

		if sortfiles.CheckTypes(fullPath, fileTypes) || fileTypes[0] == "all" {
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
			fmt.Println("--------------------------------")
		} else {
			fmt.Println("file", fullPath, "skipped")
		}

	}
}

func readString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

package sortfiles

import (
	"fmt"
	"os"
	"path"

	"github.com/rwcarlsen/goexif/exif"
)

func GetData(filePath string) (EXIFdata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return EXIFdata{}, err
	}
	defer file.Close()

	data, err := exif.Decode(file)
	if err != nil {
		return EXIFdata{}, err
	}

	taken, _ := data.DateTime()
	var exifdata EXIFdata

	exifdata.day = uint8(taken.Day())
	exifdata.month = uint8(taken.Month())
	exifdata.year = uint16(taken.Year())

	return exifdata, nil
}

func CreateFolder(exif *EXIFdata, path string) (string, error) {
	createdDir := fmt.Sprintf("%v/%v_%v_%v", path, exif.year, exif.month, exif.day)
	err := os.Mkdir(createdDir, os.ModePerm)
	return createdDir, err
}

func MoveToFolder(oldPath string, newPath string) error {
	err := os.Rename(oldPath, newPath)
	return err
}

func CheckTypes(filePath string, fileTypes []string) (status bool) {
	for i := 0; i < len(fileTypes); i++ {
		if fileTypes[i] != path.Ext(filePath) {
			status = false
		} else {
			status = true
			break
		}
	}
	return status
}

func IsItDir(filePath string) (status bool) {
	if file, _ := os.Stat(filePath); file.IsDir() {
		return true
	}
	return false
}

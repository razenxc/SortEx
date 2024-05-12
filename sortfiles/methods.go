package sortfiles

import (
	"fmt"
	"os"

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

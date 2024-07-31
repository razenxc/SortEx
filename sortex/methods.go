package sortex

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

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

func MoveToDir(oldPath string, newPath string) error {
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

func CreateBackup(pathToSave string, data []BackupJSON) error {
	cdt := getCurrentDateTime() // current data time

	file, err := os.Create(pathToSave + "/" + fmt.Sprintf("%v_%v_%v-%v_%v_%v-backup", cdt.day, cdt.month, cdt.year, cdt.hour, cdt.minute, cdt.second) + ".sfbackup")
	//file, err := os.Create(pathToSave + "/backup.sfbackup")
	if err != nil {
		return err
	}
	defer file.Close()

	toWrite, err := json.Marshal(data)
	if err != nil {
		return err
	}

	nBytes, err := file.Write(toWrite)
	if err != nil {
		return err
	}
	fmt.Println("=====================")
	fmt.Printf("| writed bytes: %v\n", nBytes)
	fmt.Println("=====================")
	return nil
}

func RevertChanges(pathToBackup string) error {
	byteData, err := os.ReadFile(pathToBackup)
	if err != nil {
		return err
	}

	var data []BackupJSON

	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		err := MoveToDir(string(data[i].NewPath), string(data[i].OldPath))
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func getCurrentDateTime() (d *dateTime) {
	time := time.Now()
	d = &dateTime{
		hour:   uint8(time.Hour()),
		minute: uint8(time.Minute()),
		second: uint8(time.Second()),
		day:    uint8(time.Day()),
		month:  uint8(time.Month()),
		year:   uint16(time.Year()),
	}
	return
}

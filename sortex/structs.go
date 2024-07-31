package sortex

type EXIFdata struct {
	Date
}

type Date struct {
	day   uint8
	month uint8
	year  uint16
}

type BackupJSON struct {
	OldPath string `json:"OldPath"`
	NewPath string `json:"NewPath"`
}

type dateTime struct {
	hour, minute, second, day, month uint8
	year                             uint16
}

package sortfiles

type EXIFdata struct {
	Date
}

type Date struct {
	day   uint8
	month uint8
	year  uint16
}

package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

type DataRecordInfo struct {
	NumberAxes                   int
	NumberBytesInGeneralBlock    int
	NumberBytesInCoordinateBlock int
	NumberBytesInAxisBlock       int
}

// Reports the diagnostic code (0) or code plus text (1) for
// the last rejected command
func (d *DMC41x3) GetErrorCode(mode int) (string, error) {
	cmd := fmt.Sprintf("TC %d", mode)
	return d.Query(cmd)
}

// Returns the internal clock in milliseconds since power-on
func (d *DMC41x3) GetTime() (int, error) {
	return d.QueryInt("MG TIME")
}

// Gets the hardware configuration and controller data structure
func (d *DMC41x3) GetDataRecordInformation() (DataRecordInfo, error) {
	response, err := d.Query("QZ")
	if err != nil {
		return DataRecordInfo{}, nil
	}
	values := strings.Split(strings.TrimSpace(response), ", ")
	if len(values) < 4 {
		return DataRecordInfo{}, fmt.Errorf("invalid response, got %s", response)
	}
	numberAxes, _ := strconv.Atoi(values[0])
	numberBytesInGB, _ := strconv.Atoi(values[1])
	numberBytesInCB, _ := strconv.Atoi(values[2])
	numberBytesInAB, _ := strconv.Atoi(values[3])

	return DataRecordInfo{
		NumberAxes:                   numberAxes,
		NumberBytesInGeneralBlock:    numberBytesInGB,
		NumberBytesInCoordinateBlock: numberBytesInCB,
		NumberBytesInAxisBlock:       numberBytesInAB,
	}, nil
}

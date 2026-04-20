package protocol

import "fmt"

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
	responseBytes := []byte(response)
	return DataRecordInfo{
		NumberAxes:                   int(responseBytes[0]),
		NumberBytesInGeneralBlock:    int(responseBytes[1]),
		NumberBytesInCoordinateBlock: int(responseBytes[2]),
		NumberBytesInAxisBlock:       int(responseBytes[3]),
	}, nil
}

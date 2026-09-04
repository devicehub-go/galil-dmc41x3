/*
Author: Leonardo Rossi Leao
Created at: April 9th, 2026
*/

package protocol

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/devicehub-go/unicomm"
)

type DMC41x3 struct {
	Communication unicomm.Unicomm
	mutex         sync.Mutex
}

var (
	numericRegex = regexp.MustCompile(`[-+]?\d*\.?\d+`)
)

// Establishes a connection with galil controller
func (d *DMC41x3) Connect() error {
	if err := d.Communication.Connect(); err != nil {
		return err
	}
	fmt.Println("Connected to galil controller")
	return nil
}

// Closes the connection with galil controller
func (d *DMC41x3) Disconnect() error {
	return d.Communication.Disconnect()
}

// Checks if there is a connection established with controller
func (d *DMC41x3) IsConnected() bool {
	return d.Communication.IsConnected()
}

// Writes a command to the controller
func (d *DMC41x3) write(command string) error {
	if !d.IsConnected() {
		return fmt.Errorf("no galil controller connected")
	}
	return d.Communication.Write([]byte(command + "\r"))
}

// Reads the response from the controller until the delimiter
func (d *DMC41x3) read() ([]byte, error) {
	if !d.IsConnected() {
		return nil, fmt.Errorf("no galil controller connected")
	}

	response := make([]byte, 0)
	for {
		b, err := d.Communication.Read(1)
		if err != nil {
			return nil, err
		}
		if len(b) == 0 {
			continue
		}
		if b[0] == ':' {
			break
		}
		if b[0] == '?' {
			if err := d.write("TC 1"); err != nil {
				return nil, errors.New("command rejected: error on request error code")
			}
			response, err := d.read()
			if err != nil {
				return nil, errors.New("command rejected: error on read error code")
			}
			return nil, fmt.Errorf("command rejected: %s", string(response))
		}
		response = append(response, b[0])
	}
	return response, nil
}

func (d *DMC41x3) WriteCommand(command string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := d.write(command); err != nil {
		return err
	}
	_, err := d.read()
	return err
}

func (d *DMC41x3) Query(command string) (string, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := d.write(command); err != nil {
		return "", err
	}
	response, err := d.read()
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// Queries a float64 number from the controller
func (d *DMC41x3) QueryFloat64(command string) (float64, error) {
	response, err := d.Query(command)
	if err != nil {
		return 0, err
	}
	match := numericRegex.FindString(strings.TrimSpace(response))
	if match == "" {
		return 0, fmt.Errorf("no numeric value found in response: %s", response)
	}
	return strconv.ParseFloat(match, 64)
}

// Queries an int number from the controller
func (d *DMC41x3) QueryInt(command string) (int, error) {
	response, err := d.Query(command)
	if err != nil {
		return 0, err
	}
	match := numericRegex.FindString(strings.TrimSpace(response))
	if match == "" {
		return 0, fmt.Errorf("no numeric value found in response: %s", response)
	}
	return strconv.Atoi(match)
}

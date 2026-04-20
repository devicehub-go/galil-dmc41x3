package protocol

import (
	"fmt"
	"strings"
)

// Downloads a program to RAM memory of controller
func (d *DMC41x3) DownloadProgram(program string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := d.write("DL"); err != nil {
		return err
	}
	for line := range strings.SplitSeq(program, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if err := d.Communication.Write([]byte(trimmed + "\r")); err != nil {
			return err
		}
	}
	if err := d.Communication.Write([]byte("\x1a")); err != nil {
		return err
	}
	_, err := d.read()
	return err
}

// Starts the execution of the program in a specific thread
func (d *DMC41x3) ExecuteProgram(label string, thread int) error {
	if thread < 0 || 7 < thread {
		return fmt.Errorf("thread must be 0-7, got %d", thread)
	}
	if !strings.HasPrefix(label, "#") {
		label = "#" + label
	}
	cmd := fmt.Sprintf("XQ%s,%d", label, thread)
	return d.WriteCommand(cmd)
}

// Stops the program execution in a specific thread
func (d *DMC41x3) StopProgram(thread int) error {
	cmd := fmt.Sprintf("HX%d", thread)
	return d.WriteCommand(cmd)
}

// Returns the code in controller memory
func (d *DMC41x3) GetProgram() (string, error) {
	return d.Query("LS")
}

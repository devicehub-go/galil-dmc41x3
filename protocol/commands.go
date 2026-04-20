package protocol

import "fmt"

// Starts the motion profile
func (d *DMC41x3) Begin(axes string) error {
	cmd := fmt.Sprintf("BG%s", axes)
	return d.WriteCommand(cmd)
}

// Commands a controlled deceleration to zero velocity
func (d *DMC41x3) Stop(axes string) error {
	cmd := fmt.Sprintf("ST%s", axes)
	return d.WriteCommand(cmd)
}

// Enables the motor amplifier
func (d *DMC41x3) Enable(axes string) error {
	cmd := fmt.Sprintf("SH%s", axes)
	return d.WriteCommand(cmd)
}

// Disables the motor amplifier
func (d *DMC41x3) Disable(axes string) error {
	cmd := fmt.Sprintf("MO%s", axes)
	return d.WriteCommand(cmd)
}

// Sets a digital output to high (ON)
func (d *DMC41x3) SetDigitalOutput(output int) error {
	if output < 1 || 16 < output {
		return fmt.Errorf("output must be 1-16, got %d", output)
	}
	cmd := fmt.Sprintf("SB%d", output)
	return d.WriteCommand(cmd)
}

// Sets a digital output to low (OFF)
func (d *DMC41x3) ClearDigitalOutput(output int) error {
	if output < 1 || 16 < output {
		return fmt.Errorf("output must be 1-16, got %d", output)
	}
	cmd := fmt.Sprintf("CB%d", output)
	return d.WriteCommand(cmd)
}

// Saves the program code to flash
func (d *DMC41x3) BurnProgram() error {
	return d.WriteCommand("BP")
}

// Saves variables and parameters
func (d *DMC41x3) Burn() error {
	return d.WriteCommand("BN")
}

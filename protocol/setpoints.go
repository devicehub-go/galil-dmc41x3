package protocol

import "fmt"

// Sets the motor type
func (d *DMC41x3) SetMotorType(axis rune, motorType float64) error {
	cmd := fmt.Sprintf("MT%c=%.1f", axis, motorType)
	return d.WriteCommand(cmd)
}

// Sets the level of current that driver deliveries to the motor
func (d *DMC41x3) SetAmplifierGain(axis rune, gain int) error {
	cmd := fmt.Sprintf("AG%c=%d", axis, gain)
	return d.WriteCommand(cmd)
}

// Sets the stepper motor behavior when it is holding
func (d *DMC41x3) SetLowCurrent(axis rune, behavior int) error {
	cmd := fmt.Sprintf("LC%c=%d", axis, behavior)
	return d.WriteCommand(cmd)
}

// Sets an absolute positon in motor counts
func (d *DMC41x3) SetTargetAbsPosition(axis rune, position int) error {
	cmd := fmt.Sprintf("PA%c=%d", axis, position)
	return d.WriteCommand(cmd)
}

// Sets a relative position in motor counts
func (d *DMC41x3) SetTargetRelPosition(axis rune, position int) error {
	cmd := fmt.Sprintf("PR%c=%d", axis, position)
	return d.WriteCommand(cmd)
}

// Sets velocity for continuous motion in counts/s
func (d *DMC41x3) SetJogSpeed(axis rune, velocity int) error {
	cmd := fmt.Sprintf("JG%c=%d", axis, velocity)
	return d.WriteCommand(cmd)
}

// Sets velocity for positioning moves in counts/s
func (d *DMC41x3) SetSpeed(axis rune, velocity int) error {
	cmd := fmt.Sprintf("SP%c=%d", axis, velocity)
	return d.WriteCommand(cmd)
}

// Sets linear acceleration rate in counts/s²
func (d *DMC41x3) SetAcceleration(axis rune, acceleration int) error {
	cmd := fmt.Sprintf("AC%c=%d", axis, acceleration)
	return d.WriteCommand(cmd)
}

// Sets linear deceleration rate in counts/s²
func (d *DMC41x3) SetDeceleration(axis rune, deceleration int) error {
	cmd := fmt.Sprintf("DC%c=%d", axis, deceleration)
	return d.WriteCommand(cmd)
}

// Sets continuos voltage limit for driver protection
func (d *DMC41x3) SetContinousVoltageLimit(axis rune, limit float64) error {
	if limit < 0 || 9.998 < limit {
		return fmt.Errorf("limit must be 0-9.998, got %.3f", limit)
	}
	cmd := fmt.Sprintf("TL%c=%.3f", axis, limit)
	return d.WriteCommand(cmd)
}

// Sets the peak voltage limit for driver protection
func (d *DMC41x3) SetPeakVoltageLimit(axis rune, limit float64) error {
	if limit < 0 || 9.998 < limit {
		return fmt.Errorf("limit must be 0-9.998, got %.3f", limit)
	}
	cmd := fmt.Sprintf("TK%c=%.3f", axis, limit)
	return d.WriteCommand(cmd)
}

// Sets the smoothing constant for step pulses.
// Higher values reduce jerk but add delay
func (d *DMC41x3) SetSmoothing(axis rune, constant float64) error {
	if constant < 0.5 || 128 < constant {
		return fmt.Errorf("constant must be 0.5-128, got %f", constant)
	}
	cmd := fmt.Sprintf("KS%c=%f", axis, constant)
	return d.WriteCommand(cmd)
}

// Defines the value of the current position
func (d *DMC41x3) DefinePosition(axis rune, position float64) error {
	cmd := fmt.Sprintf("DP%c=%f", axis, position)
	return d.WriteCommand(cmd)
}

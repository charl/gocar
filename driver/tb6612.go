package driver

/*
	TB6612 DC motor driver.

	Set direction_channel to the GPIO channel which connect to MA,
	Set motor_B to the GPIO channel which connect to MB,
	Both GPIO channel use BCM numbering;
	Set pwm_channel to the PWM channel which connect to PWMA,
	Set pwm_B to the PWM channel which connect to PWMB;
	PWM channel using PCA9685, Set pwm_address to your address, if is not 0x40
	Set debug to True to print out debug informations.
*/

import (
	"errors"

	"github.com/stianeikeland/go-rpio"
)

// Motor is a DC motor type.
type Motor struct {
	channel          int
	directionChannel int
	pin              rpio.Pin
	Offset           bool
	forwardOffset    bool
	backwardOffset   bool
	Speed            int
}

// NewMotor creates a new DC motor instance.
func NewMotor(channel, directionChannel int) *Motor {
	return &Motor{
		channel:          channel,
		directionChannel: directionChannel,
		pin:              rpio.Pin(channel),
		Offset:           true,
		forwardOffset:    true,
		backwardOffset:   false,
		Speed:            0,
	}
}

// Init the TB6612 DC motor driver controller.
func (d *Motor) Init() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	d.pin.Mode(rpio.Pwm)
	d.pin.Freq(100000)
	d.pin.DutyCycle(0, 100)
	rpio.StartPwm()

	return nil
}

// Close cleans up any used resources.
func (d *Motor) Close() error {
	return rpio.Close()
}

// SetPWM sets the duty cycle.
func (d *Motor) SetPWM(value uint32) {
	d.pin.DutyCycle(value, 100)
}

// SetSpeed sets the motor speed.
func (d *Motor) SetSpeed(speed int) error {
	if speed < 0 || speed > 100 {
		return errors.New("speed must be in 0 to 100 range")
	}
	d.Speed = speed

	return nil
}

// SetOffset sets the offset.
func (d *Motor) SetOffset(offset bool) {
	d.forwardOffset = offset
	d.backwardOffset = !d.forwardOffset
}

// Forward sets the motor direction to forward.
func (d *Motor) Forward() {
	// GPIO.output(self.direction_channel, self.forward_offset)
	// self.Speed = self._speed
}

// Backward sets the motor direction backward.
func (d *Motor) Backward() {
	// GPIO.output(self.direction_channel, self.backward_offset)
	// self.Speed = self._speed
}

// Stop stops the motor by setting the speed to 0.
func (d *Motor) Stop() error {
	err := d.SetSpeed(0)
	if err != nil {
		return err
	}

	return nil
}

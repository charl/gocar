package components

/*
	This component exposes the rear wheels as a single unit.
*/

import (
	"github.com/charl/gocar/driver"
)

const (
	rpi3BPBus      = 1
	pca9685Address = 0x40

	MotorA = 17
	MotorB = 27
	PWMA   = 4
	PWMB   = 5
)

// RearWheels is a rear wheel controller.
type RearWheels struct {
	forwardA bool
	forwardB bool
	left     *driver.Motor
	right    *driver.Motor
	pwm      *driver.PCA9685
	speed    int
}

// NewRearWheels creates a new rear wheels controller.
func NewRearWheels() (*RearWheels, error) {
	pwm, err := driver.NewPCA9685(rpi3BPBus, pca9685Address)
	if err != nil {
		return nil, err
	}
	return &RearWheels{
		forwardA: true,
		forwardB: true,
		left:     driver.NewMotor(MotorA),
		right:    driver.NewMotor(MotorB),
		pwm:      pwm,
		speed:    0,
	}, nil
}

// SetPWMX is a PWM value helper.
func (c *RearWheels) SetPWMX(pwm, value int16) error {
	pulseWidth := c.pwm.Map(value, 0, 100, 0, 4095)
	err := c.pwm.Write(pwm, 0, pulseWidth)
	if err != nil {
		return err
	}
	return nil
}

// SetPWMA sets the PWM for wheel A.
func (c *RearWheels) SetPWMA(value int16) error {
	return c.SetPWMX(PWMA, value)
}

// SetPWMB sets the PWM for wheel B.
func (c *RearWheels) SetPWMB(value int16) error {
	return c.SetPWMX(PWMB, value)
}

// SetSpeed set the speed for both wheels.
func (c *RearWheels) SetSpeed(value int) {
	c.left.Speed = c.speed
	c.right.Speed = c.speed
}

// Forward moves both wheels forward.
func (c *RearWheels) Forward() {
	c.left.Forward()
	c.right.Forward()
}

// Backward moves both wheels backwards.
func (c *RearWheels) Bacward() {
	c.left.Backward()
	c.right.Backward()
}

// Stop sets the speed for both wheels to 0.
func (c *RearWheels) Stop() {
	c.left.Stop()
	c.right.Stop()
}

// Ready sets the back wheels to the ready position.
func (c *RearWheels) Ready() {
	c.left.Offset = c.forwardA
	c.right.Offset = c.forwardB
	c.Stop()
}

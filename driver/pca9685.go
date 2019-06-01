package driver

/*
	PCA9685 driver module.
*/

import (
	"math"
	"time"

	"github.com/d2r2/go-i2c"
)

const (
	pca9685Mode1      = 0x00
	pca9685Mode2      = 0x01
	pca9685Subadr1    = 0x02
	pca9685Subadr2    = 0x03
	pca9685Subadr3    = 0x04
	pca9685Prescale   = 0xFE
	pca9685LED0OnL    = 0x06
	pca9685LED0OnH    = 0x07
	pca9685LED0OffL   = 0x08
	pca9685LED0OffH   = 0x09
	pca9685AllLEDOnL  = 0xFA
	pca9685AllLEDOnH  = 0xFB
	pca9685AllLEDOffL = 0xFC
	pca9685AllLEDOffH = 0xFD

	pca9685Restart = 0x80
	pca9685Sleep   = 0x10
	pca9685AllCall = 0x01
	pca9685Invrt   = 0x10
	pca9685Outdrv  = 0x04

	sleepMillis = 5
)

// PCA9685 is a PCA9685 driver type.
type PCA9685 struct {
	busNumber int
	addr      uint8
	freq      int
	bus       *i2c.I2C
}

// NewPCA9685 creates a new instance of the PCA9685 driver.
func NewPCA9685(busNumber int, addr uint8) (*PCA9685, error) {
	bus, err := i2c.NewI2C(addr, busNumber)
	if err != nil {
		return nil, err
	}

	driver := &PCA9685{
		busNumber: busNumber,
		addr:      addr,
		freq:      0,
		bus:       bus,
	}

	err = driver.SetFreq(60)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

// SetFreq sets the frequency.
func (d *PCA9685) SetFreq(freq int) error {
	d.freq = freq
	prescaleValue := 25000000.0
	prescaleValue /= 4096.0
	prescaleValue /= float64(freq)
	prescaleValue -= 1.0

	prescale := math.Floor(prescaleValue + 0.5)
	oldMode, err := d.read(pca9685Mode1)
	if err != nil {
		return err
	}

	newMode := (oldMode & 0x7F) | 0x10
	err = d.write(pca9685Mode1, newMode)
	if err != nil {
		return err
	}

	err = d.write(pca9685Prescale, int16(math.Floor(prescale)))
	if err != nil {
		return err
	}

	err = d.write(pca9685Mode1, oldMode)
	if err != nil {
		return err
	}
	d.milliSleep()

	err = d.write(pca9685Mode1, oldMode|0x80)
	if err != nil {
		return err
	}

	return nil
}

// Init the PCA9685 16-channel 12-bit PWM/Servo controller.
func (d *PCA9685) Init() error {
	err := d.WriteAll(0, 0)
	if err != nil {
		return err
	}
	err = d.write(pca9685Mode1, pca9685AllCall)
	if err != nil {
		return err
	}
	err = d.write(pca9685Mode2, pca9685Outdrv)
	if err != nil {
		return err
	}
	d.milliSleep()

	mode1, err := d.read(pca9685Mode1)
	if err != nil {
		return err
	}

	mode1 = mode1 & ^pca9685Sleep
	err = d.write(pca9685Mode1, mode1)
	if err != nil {
		return err
	}
	d.milliSleep()

	return nil
}

// Close cleans up all PCA9685 resources.
func (d *PCA9685) Close() error {
	return d.bus.Close()
}

// WriteAll sets the on/off values for all channels.
func (d *PCA9685) WriteAll(on, off int16) error {
	err := d.write(pca9685AllLEDOnL, on&0xFF)
	if err != nil {
		return err
	}

	err = d.write(pca9685AllLEDOnH, on>>8)
	if err != nil {
		return err
	}

	err = d.write(pca9685AllLEDOffL, off&0xFF)
	if err != nil {
		return err
	}

	err = d.write(pca9685AllLEDOffH, off>>8)
	if err != nil {
		return err
	}

	return nil
}

// Write writes data to a specific channel.
func (d *PCA9685) Write(channel, on, off int16) error {
	err := d.write(byte(pca9685LED0OnL+4*channel), on&0xFF)
	if err != nil {
		return err
	}

	err = d.write(byte(pca9685LED0OnH+4*channel), on>>8)
	if err != nil {
		return err
	}

	err = d.write(byte(pca9685LED0OffL+4*channel), off&0xFF)
	if err != nil {
		return err
	}

	err = d.write(byte(pca9685LED0OffH+4*channel), off>>8)
	if err != nil {
		return err
	}

	return nil
}

// write writes the relevant data to the I2C slave.
func (d *PCA9685) write(reg byte, value int16) error {
	return d.bus.WriteRegS16LE(reg, value)
}

// read reads the relevant data from the I2C slave.
func (d *PCA9685) read(reg byte) (int16, error) {
	return d.bus.ReadRegS16BE(reg)
}

func (d *PCA9685) milliSleep() {
	time.Sleep(sleepMillis * time.Millisecond)
}

// Map the value from one range to another.
func (d *PCA9685) Map(x, inMin, inMax, outMin, outMax int16) int16 {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}

// if __name__ == '__main__':
//     import time

//     pwm = PWM()
//     pwm.frequency = 60
//     for i in range(16):
//         time.sleep(0.5)
//         print '\nChannel %d\n' % i
//         time.sleep(0.5)
//         for j in range(4096):
//             pwm.write(i, 0, j)
//             print 'PWM value: %d' % j
//             time.sleep(0.0003)

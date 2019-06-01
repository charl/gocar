# gocar

[Go](https://golang.org/) port to control the SunFounder [PiCar-V](https://github.com/sunfounder/SunFounder_PiCar-V).

## Components

The car uses a [Raspberry PI 3B](https://www.raspberrypi.org/products/raspberry-pi-3-model-b-plus/) as the controller.

### ICs

* Robot HATS:
  * [MP1593](https://github.com/sunfounder/SunFounder_PiCar-V/blob/master/datasheet/MP1593.pdf)
  * [PCF8591](https://github.com/sunfounder/SunFounder_PiCar-V/blob/master/datasheet/PCF8591.pdf)
* [PCA9685](https://github.com/sunfounder/SunFounder_PiCar-V/blob/master/datasheet/PCA9685.pdf)
* [TB6612FNG](https://github.com/sunfounder/SunFounder_PiCar-V/blob/master/datasheet/TB6612FNG.pdf)

#### MP1593

The MP1593 is a step-down regulator with an internal Power MOSFET.

It is mounted on the SunFounder HAT attached to the Raspberry Pi.

#### PCF8591

The PCF8591 is a single-chip, single-supply low-power 8-bit CMOS data acquisition device with four analog inputs, one analog output and a serial I2C-bus interface.

Three address pins A0, A1 and A2 are used for programming the hardware address, allowing the use of up to eight devices connected to the I2C-bus without additional hardware. Address, control and data to and from the device are transferred serially via the two-line bidirectional I2C-bus.

The functions of the device include analog input multiplexing, on-chip track and hold function, 8-bit analog-to-digital conversion and an 8-bit digital-to-analog conversion. The maximum conversion rate is given by the maximum speed of the I2C-bus.

#### PCA9685

The PCA9685 is an I2C-bus controlled 16-channel LED controller optimized for Red/Green/Blue/Amber (RGBA) color backlighting applications.

Each LED output has its own 12-bit resolution (4096 steps) fixed frequency individual PWM controller that operates at a programmable frequency from a typical of 24Hz to 1526Hz with a duty cycle that is adjustable from 0 % to 100 % to allow the LED to be set to a specific brightness value.

All outputs are set to the same PWM frequency.

Each LED output can be off or on (no PWM control), or set at its individual PWM controller value. The LED output driver is programmed to be either open-drain with a 25 mA current sink capability at 5V or totem pole with a 25mA sink, 10mA source capability at 5V.

The PCA9685 operates with a supply voltage range of 2.3V to 5.5V and the inputs and outputs are 5.5V tolerant. LEDs can be directly connected to the LED output (up to 25mA, 5.5V) or controlled with external drivers and a minimum amount of discrete components for larger current or higher voltage LEDs.

The PCA9685 is in the new Fast-mode Plus (Fm+) family. Fm+ devices offer higher frequency (up to 1 MHz) and more densely populated bus operation (up to 4000 pF).

#### TB6612FNG

The TB6612FNG is a driver IC for dual DC motor with output transistor in LD MOS structure with low ON-resistor. Two input signals, IN1 and IN2, can choose one of four modes such as CW, CCW, short brake, and stop mode.

### Servo

The SunFounder [SF0180 9g](https://github.com/sunfounder/SunFounder_PiCar-V/blob/master/datasheet/SunFounder_SF0180_Servo_datasheet.pdf) servos.

### DC Motors

The kit ships with two rear-mounted DC motors that drive the rear wheel.

### Camera

It ships with a wide angle USB camera.

### Architecture

The original Python code is divided into two parts: a server and a client. Server runs under Python v2 and Django v1.9 while the client runs under Python v3 and PyQt v5.

The server runs on the Raspberry Pi and you can run the PyQT client on Windows and Linux. The Django server also exposes a web app from where you can control the car via any web browser connected to the same network as the Raspberry Pi.

The video from the USB camera is streamed to the web app using [mjpg-streamer](https://github.com/jacksonliam/mjpg-streamer).

#### Front Wheels

The front wheels are connected up to a SunFounder SF0180 9g Servo (left 90°/right 90° articulation from centre) that takes GND/VIN/PWM input from the PCA9685 I2C bus on pin 0. The VIN is supplied from the SunFounder HAT (Servo Port) that plugs into the Raspberry Pi.

#### Rear Wheels

The rear wheels are driven by DC motors that are powered via the TB6612FNG Motor Driver (Left-rear is plugged into A1/A2 and right-rear is plugged into B1/B2).

The TB6612FNG Motor Driver takes PWM input for the DC motors (PWMA (left-rear) and PWMB (right-rear)) via PWM pins 4 and 5 on the PCA9685 I2C bus.

It is powered via the SunFounder HAT (Motor Port) that plugs into the Raspberry Pi.

#### Camera Pan/Tilt

Two SunFounder SF0180 9g Servos are used to provide pan/tilt functionality to the camera. They take GND/VIN/PWM input from the PCA9685 I2C bus on pins 2 (tilt) and 1 (pan). The VIN is supplied from the SunFounder HAT (Servo Port) that plugs into the Raspberry Pi.
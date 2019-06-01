package main

import (
	"log"
	"time"

	"github.com/charl/gocar/driver"
)

const (
	rpi3BPBus      = 1
	pca9685Address = 0x40
)

func main() {
	// d, err := driver.NewPCA9685(rpi3BPBus, pca9685Address)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf(">>>> driver: %#v", d)

	// err = d.Init()
	// if err != nil {
	// 	panic(err)
	// }

	// Test the Motor.
	a := driver.NewMotor(17, 23)
	err := a.Init()
	if err != nil {
		panic(err)
	}
	defer a.Close()

	b := driver.NewMotor(27, 24)
	err = b.Init()
	if err != nil {
		panic(err)
	}
	defer b.Close()

	a.Forward()
	for i := 0; i < 100; i++ {
		a.SetSpeed(i)
		log.Printf("Forward: %d", i)
		time.Sleep(5 * time.Millisecond)
	}
	for i := 99; i >= 0; i-- {
		a.SetSpeed(i)
		log.Printf("Forward: %d", i)
		time.Sleep(5 * time.Millisecond)
	}
	a.Backward()
	for i := 0; i < 100; i++ {
		a.SetSpeed(i)
		log.Printf("Backward: %d", i)
		time.Sleep(5 * time.Millisecond)
	}
	for i := 99; i >= 0; i-- {
		a.SetSpeed(i)
		log.Printf("Backward: %d", i)
		time.Sleep(5 * time.Millisecond)
	}

	// 	motorB.forward()
	// 	for i in range(0, 101):
	// 		motorB.speed = i
	// 		time.sleep(delay)
	// 	for i in range(100, -1, -1):
	// 		motorB.speed = i
	// 		time.sleep(delay)

	// 	motorB.backward()
	// 	for i in range(0, 101):
	// 		motorB.speed = i
	// 		time.sleep(delay)
	// 	for i in range(100, -1, -1):
	// 		motorB.speed = i
	// 		time.sleep(delay)

	// if __name__ == '__main__':
	// 	test()

	// Test the rear wheels.
	// bw, err := components.NewRearWheels()
	// if err != nil {
	// 	panic(err)
	// }
	// defer bw.Stop()

	// bw.Forward()
	// for i := 0; i < 100; i++ {
	// 	bw.SetSpeed(i)
	// 	log.Printf("Forward: %d", i)
	// 	time.Sleep(5 * time.Millisecond)
	// }
	// for i := 99; i >= 0; i-- {
	// 	bw.SetSpeed(i)
	// 	log.Printf("Forward: %d", i)
	// 	time.Sleep(5 * time.Millisecond)
	// }

	// bw.Bacward()
	// for i := 0; i < 100; i++ {
	// 	bw.SetSpeed(i)
	// 	log.Printf("Backward: %d", i)
	// 	time.Sleep(5 * time.Millisecond)
	// }
	// for i := 99; i >= 0; i-- {
	// 	bw.SetSpeed(i)
	// 	log.Printf("Backward: %d", i)
	// 	time.Sleep(5 * time.Millisecond)
	// }
}
